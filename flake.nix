{
	inputs = {
		nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
		gotk4-nix.url = "github:diamondburned/gotk4-nix/main";
		gotk4-nix.flake = false;
		flake-utils.url = "github:numtide/flake-utils";
	};

	outputs =
		{
			self,
			nixpkgs,
			gotk4-nix,
			flake-utils,
		}:

		flake-utils.lib.eachDefaultSystem (system:
			let
				pkgs = nixpkgs.legacyPackages.${system};
			in
			{
				devShells.default = import "${gotk4-nix}/shell.nix" {
					base = {
						pname = "gotk4-spelling";

						buildInputs = pkgs: with pkgs; [
							libspelling
							gtksourceview3
							gtksourceview4
							gtksourceview5
						];
					};
					pkgs = import nixpkgs {
						inherit system;
						overlays = [
							(import "${gotk4-nix}/overlay.nix")
							(import "${gotk4-nix}/overlay-patchelf.nix")
							(self: super: {
								inherit (pkgs)
									go
									gopls
									gotools;
							})
						];
					};
				};
			}
		);
}
