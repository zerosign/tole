extern crate lazy_static;
extern crate serde;
extern crate serde_yaml;
extern crate url;

pub mod error;
pub mod manifest;

pub mod info {
    include!(concat!(env!("OUT_DIR"), "/build.rs"));
}
