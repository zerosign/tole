extern crate clap;
extern crate crossbeam_channel;
extern crate notify;
extern crate serde;
extern crate serde_yaml;

extern crate env_logger;
extern crate log;

use clap::{App, AppSettings, Arg, ArgGroup, SubCommand};
use crossbeam_channel::unbounded;
use notify::{RecommendedWatcher, RecursiveMode, Result, Watcher};
use std::time::Duration;
use tole::info;

#[inline]
fn manifest_options<'a>() -> Arg<'a, 'a> {
    Arg::with_name("manifest")
        .short("m")
        .long("manifest")
        .help("location of target manifest")
}

#[inline]
fn auth_options<'a>() -> Vec<Arg<'a, 'a>> {
    vec![
        Arg::with_name("certificates")
            .long("certificates")
            .help("map of list location of client & server certificates"),
        Arg::with_name("tokens")
            .long("tokens")
            .help("map of list of token values of client backends"),
        Arg::with_name("tokens-intervals")
            .long("tokens-intervals")
            .help("map of list of token intervals lifecycle (if any) of client backends"),
        Arg::with_name("simple-auth")
            .long("simple-auth")
            .help("map of list of simple auth backend=[user:password] of client backends"),
    ]
}

#[inline]
fn file_creation_flags<'a>() -> ArgGroup<'a> {
    ArgGroup::with_name("file_creation_flags")
        .args(&["patch-only", "truncate", "auto-create"])
        .required(true)
}

fn main() -> Result<()> {
    //
    // [auth-options]
    // --certificates=sample=file://...,[]
    // --tokens=sample=das98djau9sd8a9s9da9,...[]
    // --tokens-interval=sample=20s,...[]
    // --simple-auth=sample=dasdasdaaawdaw,...[]
    //
    // [file-creation-flags]
    // --patch-only
    // --truncate
    // --auto-create
    //
    // tole lint --manifest [manifest.yml]
    // tole check [auth-options] --manifest [manifest.yml]
    // tole eval [auth-options] --manifest [manifest.yml]
    // tole watch [auth-options] [file-creation-flags] --manifest [manifest.yml]
    //
    let matches = App::new("tole")
        .version(info::PKG_VERSION)
        .setting(AppSettings::GlobalVersion)
        .author("zerosign <r1nlx0@gmail.com>")
        .about("configuration & secrets management for microservice era")
        .arg(Arg::with_name("verbose")
             .short("v")
             .long("verbose")
             .help("set verbosity level")
        )
        .subcommands(vec![
            SubCommand::with_name("lint")
                .about("check for manifest & template correctness (without evaluating or touching remote sources or output)")
                .arg(manifest_options()),
            SubCommand::with_name("check")
                .about("check for source readiness")
                .arg(manifest_options())
                .args(&auth_options()),
            SubCommand::with_name("eval")
                .about("eval a template configuration once")
                .arg(manifest_options())
                .args(&auth_options()),
            SubCommand::with_name("watch")
                .about("watch sources and apply any changes related to its variables into output based on manifest")
                .arg(manifest_options())
                .arg(Arg::with_name("dev-mode").help("don't listen to http interface while watching").conflicts_with("listen"))
                .arg(Arg::with_name("listen").help("listen to specific socket address (for health checks etc)").conflicts_with("dev-mode"))
                .args(&auth_options())
                .group(file_creation_flags())
        ]).get_matches();

    println!("matches: {:?}", matches);

    match matches.subcommand_name() {
        Some("lint") => {
            // try convert manifest
            // show manifest errors
        }
        Some("check") => {
            println!("check");

            // convert manifest
            // initialize all sources
            // ping all sources
            // check for outputs availability
        }
        Some("eval") => {
            println!("eval");

            // convert manifest
            // initialize engine (initialize sources)
            // check for outputs availability
            // lookup template files once
            // build ast tree
            // fetch all values from the tree
            // eval the tree based on the values
        }
        Some("watch") => println!("watch"),
        _ => println!("unsupported"),
    }

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
