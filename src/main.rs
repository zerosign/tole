extern crate clap;
extern crate crossbeam_channel;
extern crate notify;
extern crate serde;
extern crate serde_yaml;

use clap::App;
use crossbeam_channel::unbounded;
use notify::{RecommendedWatcher, RecursiveMode, Result, Watcher};
use std::time::Duration;

const TOLE_VERSION: &'static str = env!("APP_VERSION");
const AUTHOR: &'static str = env!("AUTHOR");

fn main() -> Result<()> {
    let matches = App::new("tole")
        .version(TOLE_VERSION)
        .author(AUTHOR)
        .about("configuration & secrets management for microservice era")
        .arg(Arg::with_name());

    // let (tx, rx) = unbounded();

    // let mut watcher: RecommendedWatcher = Watcher::new(tx, Duration::from_secs(1))?;

    // watcher.watch(".", RecursiveMode::Recursive)?;

    // loop {
    //     match rx.recv() {
    //         Ok(event) => println!("events: {:?}", event),
    //         Err(err) => println!("error: {:?}", err),
    //     }
    // }

    Ok(())
}
