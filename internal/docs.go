// Package internal should contain everything that it is internal to your repo
// and thus should not be accessible by other repos.
//
// Picture yourself depending on another repo. What would be the use of its core
// internal modules for you if the API is well defined?
//
// A good rule of thumb is to start by having all your non-main packages in here
// and then move them out to the root of the repo if they are needed by other
// repos.
package internal
