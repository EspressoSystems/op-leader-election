{ lib, stdenv, fetchFromGitHub, rustPlatform }:

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
    # outputHashes = {
    #   "ethers-2.0.8" = "sha256-03vXjQymGuzL6RdihZkFf12Rlfygq3TlFwFgMTaJp7w=";
    #   "revm-3.3.0" = "sha256-bj3Tdita5UEdHEzLBzMX8S6GcyzB53p+GWq5+o/YhqI=";
    # };
  };

  # Enable svm-rs to build without network access.
  env =
    let
      info =
        if stdenv.isLinux then {
          platform = "linux-amd64";
          sha256 = "sha256:0qn4nlqd4yyxfjaywkfxsrrskyn6f2krvkp0cy98rha3m60b7ijf";
        } else {
          platform = "macosx-amd64";
          sha256 = lib.fakeHash;
        };
      src = builtins.fetchurl {
        url = "https://binaries.soliditylang.org/${info.platform}/list.json";
        inherit (info) sha256;
      };
    in
    {
      SVM_RELEASES_LIST_JSON = "${src}";
    };

  meta = with lib; {
    description = "A blazing fast, portable and modular toolkit for Ethereum application development";
    homepage = "https://github.com/foundry-rs/foundry";
    license = licenses.unlicense;
    maintainers = [ ];
  };
}
