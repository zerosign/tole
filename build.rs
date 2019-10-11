extern crate built;

use std::{env, path};

fn main() {
    let mut options = built::Options::default();
    options.set_dependencies(true);

    let src = env::var("CARGO_MANIFEST_DIR").unwrap();
    let dst = path::Path::new(&env::var("OUT_DIR").unwrap()).join("build.rs");

    built::write_built_file_with_opts(&options, &src, &dst)
        .expect("fails to acquire build time information");
}
