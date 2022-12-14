#!/usr/bin/env perl

use v5.36.0;
use warnings;

open my $fh, ">", "zip_gen.go" or die "failed to open: $!";

print $fh <<'END';
// Code generated by generate_zip.pl; DO NOT EDIT.

package result

import (
    "github.com/shogo82148/hi/tuple"
)
END

sub mapJoin :prototype(&$) ($f, $n) {
    return join ", ", map { $f->($_) } 1..$n;
}

for my $n(2..16) {
    say $fh "";
    my $types = mapJoin { "T$_" } $n;
    my $args = mapJoin { "o$_ Result[T$_]" } $n;

    say $fh "// Zip$n returns an optional of valid $n-tuples, if all of arguments are valid.";
    say $fh "// Otherwise, one or more arguments are not valid, it returns none";
    say $fh "func Zip${n}[$types any]($args) Result[tuple.Tuple${n}[$types]] {";
    say $fh " if (" . (join " && ", map { "o$_.err == nil" } 1..$n) . ") {";
    say $fh "  return Result[tuple.Tuple${n}[$types]]{value: tuple.Tuple${n}[$types]{" . (mapJoin { "V$_: o$_.value" } $n) . "}}";
    say $fh " } else {";
    say $fh "  errs := make([]error, 0, $n)";
    for my $i (1..$n) {
        say $fh "  if o$i.err != nil {";
        say $fh "    errs = append(errs, o$i.err)";
        say $fh "  }";
    }
    say $fh "  return Result[tuple.Tuple${n}[$types]]{";
    say $fh "    err: &wrapErrors{msg: \"result: errors\", errs: errs},";
    say $fh "  }";
    say $fh " }";
    say $fh "}";
}

system("gofmt -s -w zip_gen.go");

close($fh);

1;
