# Pinki

Pinki helps developers ship software with authenticity.

Use it anywhere you would use `gpg` to sign and verify things.

## Features

 * Easy to use
 * Portable, standalone binary
 * Anonymous -- a key is just a key, nothing more
 * Doesn't touch your filesystem
 * Reads and writes standard PEM-wrapped ASN.1 (compatible with X.509, GPG)

## Installing

### Precompiled Binaries

Visit [Releases](https://releases.twuni.dev/pinki/latest/) to download a precompiled binary for your system.

#### Verifying Binaries

All releases are signed using the following [release signing key](https://releases.twuni.dev/verify.pem):

```pem
-----BEGIN PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEdDOWMNxI5f88Yck8WNcPsxDOwMbzoU/Y
cZhfoR+gwGi0wRoSscWA1xy1BQTG6PNrQlvLJbfm2vAIAImnyMmhoKS3hwcO6F+5
4QjLZQJAQHZ6G7c842gYRSnwLLQ2GIvj
-----END PUBLIC KEY-----
```

Each release binary has a corresponding file with an **.asc** prefix containing the signature for that file.

Here's an example of how you could use Pinki to verify itself.

> :bulb: In practice, you probably want to use another tool to verify Pinki itself the first time you download
> it. Once you have a genuine copy of `pinki`, then you can use it to verify updates to itself.

```sh
# Download Pinki for Linux (64-bit)
$ curl -sSL -o pinki https://releases.twuni.dev/pinki/latest/linux-amd64/pinki

# Use Pinki to verify itself
$ ./pinki verify "$(curl -sSL https://releases.twuni.dev/verify.pem)" "$(curl -sSL https://releases.twuni.dev/pinki/latest/linux-amd64/pinki.asc)" < pinki
```

If you get an output of `OK`, the signature is valid.

### Building from source

Already have `go`? Clone this repo and run `go build`.

## Usage

Pinki is designed to make it easy for you to do one of two things:

 * **Sign** your software so other people can verify its authenticity, or

 * **Verify** the authenticity of software you are using when the developers
   are using Pinki.

### Signing your software with Pinki

First, you'll need a private key. To create a new key with the
recommended (default) options:

```sh
$ pinki key create
-----BEGIN PRIVATE KEY-----
...............................................................
...............................................................
...............................................................
-----END PRIVATE KEY-----
```

Save the output somewhere safe. Put it in your password manager,
vault, or whatever you are using to keep sensitive information
safe.

Once you have a private key, you will need to *export* that in a
way that is safe for people to verify your signatures:

```sh
$ pinki key export < /path/to/your-pinki-private-key
-----BEGIN PUBLIC KEY-----
...............................................................
...............................................................
-----END PUBLIC KEY-----
```

Publish this public key somewhere that anyone you want to be able
to verify your signatures is able to access it. You can commit it
to your source code repo, publish it to your website, etc.

> :bulb: The **public key** is not sensitive! You can safely share
> it with anyone.

Now that you have a private key, you're ready to sign your first thing!

```sh
$ pinki sign "$(cat /path/to/your-pinki-private-key)" < /path/to/your-thing-1.2.3.tar.gz
-----BEGIN SIGNATURE-----
...............................................................
...............................................................
-----END SIGNATURE-----
```

Publish that signature any way you like. Conventionally, you might want to
publish it as a file with the same name as the thing you've signed, but with
a **.sig** suffix. So **foo-1.0.tgz** would have its signature in
**foo-1.0.tgz.sig**. The choice is up to you.

### Verifying a signature with Pinki

To verify a signature, you'll need three things:

 * The thing that was signed (e.g: **foo-1.2.3.tgz**)
 * The signature (e.g: **foo-1.2.3.tgz.asc**)
 * The public key of the signer (e.g: **foomaker-signing-key.pem**)

Check the release notes or installation/verification documentation of the
thing you're trying to verify for more details on where to find these things.

Once you have them, here's how you verify the thing is authentic!

```sh
$ pinki verify "$(cat /path/to/signing-key)" "$(cat /path/to/signature)" < /path/to/thing-that-was-signed
OK
```

The command will exit with status code 0 and print "OK" on success.
Otherwise, it will exit with status code 1 and print an error message.
