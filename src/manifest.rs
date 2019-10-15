use crate::error::{ManifestError, OwnershipError};
use lazy_static::lazy_static;
use serde::Deserialize;
use std::{
    collections::{HashMap, HashSet},
    convert::{Infallible, TryFrom},
    fmt,
    fs::{self, Permissions as Perm},
    io,
    iter::FromIterator,
    os::unix::fs::PermissionsExt,
    path::{Path, PathBuf},
};
use url::{form_urlencoded::Parse as QueryPairs, Url};

lazy_static! {
    // registered locals protocols/schemes
    //
    static ref LOCALS: HashSet<&'static str> = HashSet::from_iter(
        vec![
            "file+rel",
            "file",
            "dotenv+rel",
            "dir",
            "dir+rel",
            "glob+rel",
            "glob",
            "dir",
            "dir+rel"
        ]
        .into_iter()
    );
}

#[derive(Debug, Deserialize, PartialEq)]
pub enum ManifestKind {
    ConfigManifest,
}

#[derive(Debug, Deserialize)]
pub struct Mount {
    source: String,
    target: String,
}

pub struct Ownership {
    uid: Option<usize>,
    gid: Option<usize>,
    mode: Option<Perm>,
}

impl Ownership {
    #[inline]
    pub fn uid(&self) -> Option<usize> {
        self.uid
    }

    #[inline]
    pub fn gid(&self) -> Option<usize> {
        self.gid
    }

    #[inline]
    pub fn mode(&self) -> Option<&Perm> {
        self.mode.as_ref()
    }

    #[inline]
    pub fn root_only() -> Ownership {
        Ownership {
            uid: Some(0),
            gid: Some(0),
            mode: Some(Perm::from_mode(0o1700)),
        }
    }
}

impl<'a> TryFrom<QueryPairs<'a>> for Ownership {
    type Error = OwnershipError;

    #[inline]
    fn try_from(pairs: QueryPairs) -> Result<Self, Self::Error> {
        let mut uid: Option<usize> = None;
        let mut gid: Option<usize> = None;
        let mut mode: Option<Perm> = None;

        for (key, value) in pairs {
            match key.as_ref() {
                "uid" => {
                    uid = Some(
                        value
                            .parse::<usize>()
                            .map_err(|_| OwnershipError::ParseError)?,
                    );
                }
                "gid" => {
                    gid = Some(
                        value
                            .parse::<usize>()
                            .map_err(|_| OwnershipError::ParseError)?,
                    );
                }
                "mode" => {
                    mode = Some(Perm::from_mode(
                        value
                            .parse::<u32>()
                            .map_err(|_| OwnershipError::ParseError)?,
                    ));
                }
                _ => {
                    // warning that key doesn't exists
                }
            }
        }

        Ok(Ownership {
            uid: uid,
            gid: gid,
            mode: mode,
        })
    }
}

//
// file+rel:///config/database.yml?uid=<uid>&gid=<gid>&mode=<mode>
//
pub struct LocalRef {
    inner: PathBuf,
    owner: Ownership,
    recursive: bool,
}

impl LocalRef {
    #[inline]
    fn is_relative(url: &Url) -> bool {
        url.scheme().ends_with("+rel")
    }
}

pub struct RemoteRef {
    inner: PathBuf,
    options: HashMap<String, String>,
    raw: Url,
    recursive: bool,
}

impl fmt::Debug for LocalRef {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "FileRef {{ uid: {:?}, gid: {:?}, mode: {:?} }}",
            self.owner.uid, self.owner.gid, self.owner.mode
        )
    }
}

impl fmt::Debug for RemoteRef {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "CustomRef {{ raw: {} }}", self.raw)
    }
}

//
// Repr type for sum types of mount references.
//
#[derive(Debug)]
pub enum PathRef {
    LocalRef(LocalRef),
    RemoteRef(RemoteRef),
}

impl PathRef {
    #[inline]
    pub fn is_local(url: &Url) -> bool {
        LOCALS.contains(url.scheme())
    }

    #[inline]
    fn is_recursive(url: &Url) -> bool {
        *url.query_pairs()
            .filter(|(k, _)| k == "recursive")
            .map(|(_, v)| v.parse::<bool>().unwrap_or(false))
            .collect::<Vec<bool>>()
            .first()
            .unwrap_or(&false)
    }
}

impl TryFrom<Url> for LocalRef {
    type Error = ManifestError;

    #[inline]
    fn try_from(url: Url) -> Result<Self, Self::Error> {
        let path = if LocalRef::is_relative(&url) {
            // create absolute path from current relative path using
            // fs::canonicalize
            fs::canonicalize(Path::new(&format!("./{}", url.path())))
                .map_err(Self::Error::IoError)?
                .into()
        } else {
            PathBuf::from(url.path())
        };

        // fetch ownership from url
        let ownership =
            Ownership::try_from(url.query_pairs()).map_err(ManifestError::OwnershipError)?;

        Ok(LocalRef {
            inner: path.into(),
            owner: ownership,
            recursive: PathRef::is_recursive(&url),
        })
    }
}

impl TryFrom<Url> for RemoteRef {
    type Error = Infallible;

    #[inline]
    fn try_from(url: Url) -> Result<Self, Self::Error> {
        // let options = HashMap::from_iter(url.query_pairs());
        let options = HashMap::from_iter(url.query_pairs().into_owned());

        let path = Path::new(url.path());
        let is_recursive = PathRef::is_recursive(&url);

        Ok(RemoteRef {
            inner: path.into(),
            options: options,
            raw: url.clone(),
            recursive: is_recursive,
        })
    }
}

impl TryFrom<Url> for PathRef {
    type Error = ManifestError;

    #[inline]
    fn try_from(url: Url) -> Result<Self, Self::Error> {
        // check whether mount-ref is path or url
        if Self::is_local(&url) {
            LocalRef::try_from(url).map(PathRef::LocalRef)
        } else {
            RemoteRef::try_from(url)
                .map_err(|e| ManifestError::OwnershipError(OwnershipError::Infallible(e)))
                .map(PathRef::RemoteRef)
        }
    }
}

#[derive(Debug, Deserialize)]
pub struct Manifest {
    kind: ManifestKind,
    version: String,
    sources: HashMap<String, String>,
    aliases: HashMap<String, String>,
    mounts: Vec<Mount>,
}

impl Manifest {
    #[inline]
    pub fn from_reader<R>(r: R) -> Result<Manifest, ManifestError>
    where
        R: io::Read,
    {
        serde_yaml::from_reader(r).map_err(ManifestError::SerdeError)
    }
}

#[cfg(test)]
mod test {

    use crate::manifest::{Manifest, PathRef};
    use std::{
        convert::TryFrom,
        fs,
        io::{BufRead, BufReader, Read},
        path::{Path, PathBuf},
    };
    use url::Url;

    #[test]
    fn test_manifest_deserialize() {
        let path: PathBuf = {
            let mut p = PathBuf::from(env!("CARGO_MANIFEST_DIR"));
            p.push("tests/sample_manifest.yml");
            p.into()
        };

        let f = fs::OpenOptions::new()
            .read(true)
            .write(false)
            .create(false)
            .open(path)
            .expect("can't open path");

        let result: serde_yaml::Result<Manifest> = serde_yaml::from_reader(f);

        assert!(result.is_ok());
    }

    #[test]
    fn test_url_parsing_relative() {
        println!("current_dir: {:?}", std::env::current_dir());

        let relative_url = format!(
            "file+rel:///{}?uid=1&gid=20&mode0444",
            "tests/sample_manifest.yml",
        );

        let parsed = Url::parse(&relative_url).expect("can't parse the url");

        let result = PathRef::try_from(parsed);

        println!("result: {:?}", result);

        assert!(result.is_ok());
    }

    #[test]
    fn test_url_parsing() {
        let path: PathBuf = {
            let mut p = PathBuf::from(env!("CARGO_MANIFEST_DIR"));
            p.push("tests/sample_manifest.yml");
            p.into()
        };

        let sample_url = format!(
            "file:///{}?uid=1&gid=20&mode=0444",
            path.to_str().unwrap(),
        );

        let parsed = Url::parse(&sample_url).expect("can't parse the url");
        let result = PathRef::try_from(parsed);

        println!("result: {:?}", result);

        assert!(result.is_ok());
    }
}
