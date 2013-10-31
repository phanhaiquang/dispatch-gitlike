#!/usr/bin/env perl
use Modern::Perl;
use lib 'lib';
use Dispatch::GitLike;

exit Dispatch::GitLike->new->run(@ARGV) unless caller
