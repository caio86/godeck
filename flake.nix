{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      systems = [
        "x86_64-linux"
        "aarch64-linux"
      ];
      forAllSystems =
        function: nixpkgs.lib.genAttrs systems (system: function (nixpkgs.legacyPackages.${system}));
    in
    {
      packages = forAllSystems (
        pkgs:
        let
          inherit (pkgs) buildGoModule;
        in
        {
          godeck = buildGoModule {
            pname = "godeck";
            version = "1.0.0";

            src = ./.;

            vendorHash = null;

            subPackages = [ "cmd/godeck" ];
          };
        }
      );
      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            nixfmt-rfc-style
            go
            gopls
            gofumpt
            gotools
          ];
        };
      });
    };
}
