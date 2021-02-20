{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    docutils
    go
    mage
    python3Packages.pygments
    rnix-lsp
  ];
}
