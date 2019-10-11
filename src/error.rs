use std::{convert::Infallible, io};

#[derive(Debug)]
pub enum OwnershipError {
    ParseError,
    Infallible(Infallible),
    OptionDuplError,
}

#[derive(Debug)]
pub enum ManifestError {
    // OwnershipError(OwnershipError),
    OwnershipError(OwnershipError),
    IoError(io::Error),
    SerdeError(serde_yaml::Error),
}
