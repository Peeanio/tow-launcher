# tow-launcher

tow-launcher is a cli tool to work with the ruleset of Fistful of Tows 3.

## Usage

### weapon

The `weapon` subcommand works with the weapon object type in the database, from any type used. This turned out to be missiles (AA/AT), guns, and indirect fire weapons. Since these are typically loaded onto different units, they are there own type, with the idea that unit types would contain these sub-objects.
