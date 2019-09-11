use std::{convert::Infallible, str::FromStr};

#[derive(Debug)]
pub enum OwnershipError {
    ParseError(FromStr::Err),
    Infallible(Infallible),
    OptionDuplError,
}

#[derive(Debug)]
pub enum ManifestError {
    OwnershipError(OwnershipError),
}
