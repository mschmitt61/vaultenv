# vaultenv
A simple CLI tool written in Go that allows you to integrate vault paths in your .env files and automatically pull the values from those vault paths into a usable .env file.

Big thanks to [respondcreate](https://github.com/respondcreate) for the support and feedback!

# Setup
First, setup the tap and install via brew

```
brew tap mschmitt61/mschmitt61

brew install mschmitt61/mschmitt61/vaultenv

# To update the package
brew update 
brew upgrade mschmitt61/mschmitt61/vaultenv
```

You'll also need a couple of environment variables to make use of this library
   1. `VAULT_ADDR`
   2. `VAULT_ROLE`
   3. `VAULT_SECRET`

# Commands
If your application is environment variable based and you leverage vault for secrets, you can easily leverage a custom vault path format shown below in your `.env` files. You can of course also have non-vault based environment variables in the same file.

## generate

Say you have a template env file like this, with the custom vault formatting mentioned above

`.env.template.dev`
```
env=dev
username=vault::your/vault/path::devuser
password=vault::your/vault/path::devpassword
```

Note that the vault path does *not* include the `secret/` prefix

Run these commands to pull from vault and generate a `.env.dev` file!
```
vaultenv generate .env.template.dev .env.dev
```

Now the `.env.dev` will look like below

```
env=dev
username=realuser
password=realpassword
```

Easy!

## testing
```
cd cmd #or cd helpers
go test -v -coverprofile cover.out
go tool cover -html=cover.out
```
