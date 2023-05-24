#!/bin/bash

echo " === Test app build === "
task app:build app:test

echo " === Test infra (with aws connection) === "
task infra:list infra:diff

echo " === Integration Test infra (with aws connection) === "
cdkstat
task infra:test infra:deploy
cdkstat hello
task infra:destroy
cdkstat
