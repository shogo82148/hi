#!/usr/bin/env perl

use v5.36.0;
use warnings;

open my $fh, ">", "zip_gen.go" or die "failed to open: $!";

print $fh <<'END';
// Code generated by generate_zip.pl; DO NOT EDIT.

package hi
END

for my $n(2..16) {
    say $fh "";
    my $types = join ", ", map { "T$_" } 1..$n;
    my $slice_args = join ", ", map { "s$_ []T$_" } 1..$n;
    my $slices = join ", ", map { "s${_}[i]" } 1..$n;
    say $fh "// Zip$n returns a slice of $n-tuples.";
    say $fh "// The returned slice have the length of the shortest slice.";
    say $fh "func Zip${n}[S ~[]Tuple${n}[$types], $types any]($slice_args) S {";
    say $fh "	l := len(s1)";
    for my $i(2..$n) {
        say $fh "	if len(s$i) < l {";
        say $fh "		l = len(s$i)";
        say $fh "	}";
    }
    say $fh "	ret := make(S, l)";
    say $fh "	for i := 0; i < l; i++ {";
    say $fh "		ret[i] = Tuple${n}[$types]{", join(", ", map { "s${_}[i]" } 1..$n), "}";
    say $fh "	}";
    say $fh "	return ret";
    say $fh "}";
}

system("gofmt -s -w zip_gen.go");

close($fh);
