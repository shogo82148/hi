#!/usr/bin/env perl

use v5.36.0;
use warnings;

open my $fh, ">", "unzip_gen.go" or die "failed to open: $!";

print $fh <<'END';
// Code generated by generate_zip.pl; DO NOT EDIT.

package optional

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

    say $fh "// Unzip$n converts an optional of $n-tuple to optionals of each elements.";
    say $fh "func Unzip${n}[$types any](o Optional[tuple.Tuple${n}[$types]]) (" . (mapJoin { "Optional[T$_]" } $n) . ") {";
    say $fh "  return " . (mapJoin { "Optional[T$_]{valid: o.valid, value: o.value.V$_}" } $n);
    say $fh "}";
}

system("gofmt -s -w unzip_gen.go");

close($fh);
