{ lib
, stdenv
, fetchFromGitHub
, rustPlatform
, rustc
, libusb1
, darwin
}:

rustPlatform.buildRustPackage rec {
  pname = "foundry";
  version = "0.2.0-dev";

  src = fetchFromGitHub {
    owner = "foundry-rs";
    repo = pname;
    rev = "3fec8c1ca1b2dcb0497469ddfbc21713815b37c8";
    hash = "sha256-mcxaPPlxoig4nRVBVRs9FHflGSPMP9Ux+87yrP3EcDc=";
  };

  cargoLock = {
    lockFile = src + "/Cargo.lock";
    allowBuiltinFetchGit = true;
  };

  nativeBuildInputs = [
    libusb1
  ] ++ lib.optionals stdenv.isDarwin [ darwin.DarwinTools ];

  buildInputs = lib.optionals stdenv.isDarwin [ darwin.apple_sdk.frameworks.AppKit ];

  # Tests fail
  doCheck = false;

  env =
    let
      info =
        if stdenv.isLinux then {
          platform = "linux-amd64";
          sha256 = "sha256:0qn4nlqd4yyxfjaywkfxsrrskyn6f2krvkp0cy98rha3m60b7ijf";
        } else {
          platform = "macosx-amd64";
          sha256 = "sha256:0bcq98dn79gjgrbmhwy6klb7vldx7bhgm896j6kmz33msa1xn5p6";
        };
      list = builtins.fetchurl {
        url = "https://binaries.soliditylang.org/${info.platform}/list.json";
        inherit (info) sha256;
      };
    in
    {
      # Enable svm-rs to build without network access.
      SVM_RELEASES_LIST_JSON = "${list}";
      # Make `vergen` produce a meaningful version.
      VERGEN_BUILD_TIMESTAMP = "0";
      VERGEN_BUILD_SEMVER = version;
      VERGEN_GIT_SHA = src.rev;
      VERGEN_GIT_COMMIT_TIMESTAMP = "0";
      VERGEN_GIT_BRANCH = "master";
      VERGEN_RUSTC_SEMVER = rustc.version;
      VERGEN_RUSTC_CHANNEL = "stable";
      VERGEN_CARGO_PROFILE = "release";
    };

  meta = with lib; {
    description = "A blazing fast, portable and modular toolkit for Ethereum application development";
    homepage = "https://github.com/foundry-rs/foundry";
    license = with licenses; [ mit apsl20 ];
    maintainers = [ ];
  };
}
