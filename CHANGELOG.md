## [Talos System Extensions 1.12.0-alpha.2](https://github.com/siderolabs/extensions/releases/tag/v1.12.0-alpha.2) (2025-10-27)

Welcome to the v1.12.0-alpha.2 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.11/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

NVIDIA LTS: 580.95.05
NVIDIA Production: 570.195.03
NVIDIA Container Toolkit: 1.17.9
crun: 1.24
drbd: 9.2.15
ecr-credential-provider: 1.34.1
fuse3: 3.17.4
glibc: 2.42
gvisor: 20251020.0
hailort: 4.23.0
kata-containers: 3.21.0
lldpd: 1.0.20
cloudflared: 2025.10.0
nebula: 1.9.7
netbird: 0.59.8
newt: 1.5.2
nut-client: 2.8.4
qemu: 10.1.2
spin: 0.21.0
stargz-snapshotter: 0.18.0
tailscale: 1.90.2
talos-vmtoolsd: 1.4.0
youki: 0.5.6
zerotier: 1.16.0
zfs: 2.4.0-rc2
linux-firmware: 20251021
tenstorrent: 2.4.1
mdadm: 4.4

wolfi-base: sha256:815b27b8a70713c70404e44a718eddd52ea6f4a2bfad5f56455b52cd2789a9b2


### Contributors

* Andrey Smirnov
* Dmitrii Sharshakov
* Mateusz Urbanek
* Noel Georgi
* Jorik Jonker
* Steven Kreitzer
* Michael Smith
* Breland Miley
* Fred Heinecke
* Guillaume LEGRAIN
* Hugo Meyronneinc
* Max Makarov
* Robin Elfrink
* Sammy ETUR
* Steve Francis
* Thibault VINCENT
* veymax
* Łukasz Szczepański

### Changes
<details><summary>44 commits</summary>
<p>

* [`fc85294`](https://github.com/siderolabs/extensions/commit/fc85294ac92a7fb3ee6cf99b737e739c63696afb) feat: bump deps
* [`ed8c228`](https://github.com/siderolabs/extensions/commit/ed8c228c4fea1b4515dd5d286798cfcfec3bdb47) feat: update cloudflared to 2025.10.0
* [`0beff58`](https://github.com/siderolabs/extensions/commit/0beff582fc439e159a83a11b76ddef74dc344b4a) feat: update dependencies
* [`9d782c5`](https://github.com/siderolabs/extensions/commit/9d782c572a2fcdc086fecd5aded326772ab78b40) feat: fabricmanager support Blackwell baseboards (DGX/HGX B100/B200/B300)
* [`823659b`](https://github.com/siderolabs/extensions/commit/823659bcaea94a55c54a912b84075777070d50f1) feat: use `image-signer`
* [`3f0e872`](https://github.com/siderolabs/extensions/commit/3f0e8727f353b60ffe8ad651a3804c25ad12e23f) feat: add xe extension
* [`e38736c`](https://github.com/siderolabs/extensions/commit/e38736c2ba56efde6d862e39f6b0b786ec5ba5fa) chore: update external version references
* [`38f9757`](https://github.com/siderolabs/extensions/commit/38f975746c903d10df35ce074c8713deb4e77345) chore: update dependencies
* [`83ee3b0`](https://github.com/siderolabs/extensions/commit/83ee3b07c140e8fe2020a293b5cf4e4c297476ce) feat: update tools and pkgs for new Go and NVIDIA LTS
* [`563da22`](https://github.com/siderolabs/extensions/commit/563da22175e2ba46aab11b9ef47318f9c1caa655) feat: update dependencies to match pkgs
* [`5133393`](https://github.com/siderolabs/extensions/commit/5133393ba9a2e210cd8ed920e8b121aa4022763f) fix: try fixing the grype scan
* [`e0946ad`](https://github.com/siderolabs/extensions/commit/e0946ad4b4879e53503636d2cb091a3cfb17139f) release(v1.12.0-alpha.1): prepare release
* [`6962487`](https://github.com/siderolabs/extensions/commit/6962487685afd667999d6bfc175b646ff9366aba) feat: update extensions
* [`f106fb4`](https://github.com/siderolabs/extensions/commit/f106fb4ac3d0d3db8055938a60c71e799b892f84) feat: add usb-audio-drivers
* [`b382dba`](https://github.com/siderolabs/extensions/commit/b382dba8c747124b4b3985ef2f3d279af155c83d) feat: add netbird
* [`17be2f7`](https://github.com/siderolabs/extensions/commit/17be2f7b15e9c17e21f7fa2d77d250dbd2a0e68e) chore: update linux-firmware
* [`12654ec`](https://github.com/siderolabs/extensions/commit/12654ec141789239fd3aed9f09dc427b269206ad) feat: update vmtoolsd to v1.4.0
* [`47051c6`](https://github.com/siderolabs/extensions/commit/47051c6d5da0ab9daad21f158c54bfbea8817347) feat: update NVIDIA LTS to 580.82.07
* [`caa10a4`](https://github.com/siderolabs/extensions/commit/caa10a495fec59f2a726bf8df728ead8ec74e186) docs: update zfs readme for existing pools
* [`b86a64b`](https://github.com/siderolabs/extensions/commit/b86a64b230e2b77d8850dcb2393b5095dc018c01) feat: update dependencies
* [`ea43af9`](https://github.com/siderolabs/extensions/commit/ea43af91e0d9c730742821b8f58ffeebe9b003ef) feat: update Go to 1.25.1
* [`ffcd027`](https://github.com/siderolabs/extensions/commit/ffcd027e34114cb90bedb76eb53d71c6989b3852) chore: redo dependencies as SVG
* [`ef73a85`](https://github.com/siderolabs/extensions/commit/ef73a85d368405763143382c6e66dfe78e7cbf55) feat: update dependencies
* [`2273c9b`](https://github.com/siderolabs/extensions/commit/2273c9be9c4522499a476d46e3ab239991cb766e) release(v1.12.0-alpha.0): prepare release
* [`8976af3`](https://github.com/siderolabs/extensions/commit/8976af37b40f38b2135847fb88047ec44bd19203) fix: nvidia-container-toolkit: use SHA512 from a variable for tirpc
* [`5de79a6`](https://github.com/siderolabs/extensions/commit/5de79a6ba36aa306bb6145807169d41b23eb2c8b) feat: update to Linux 6.16 and GCC 15
* [`d6bc2ce`](https://github.com/siderolabs/extensions/commit/d6bc2ce44b67a6545e235ab733d20346f11fdc71) docs: fix catalog for qemu-guest-agent
* [`5ccda4b`](https://github.com/siderolabs/extensions/commit/5ccda4bdeb79fb916b52122dcc16df6444984735) feat: zerotier - add possible custom planet file
* [`df7a67d`](https://github.com/siderolabs/extensions/commit/df7a67dc8d985f2b104caf524efc530ea0bc5634) chore: refactor manifest for talos-vmtoolsd
* [`d7a938f`](https://github.com/siderolabs/extensions/commit/d7a938fe06b34247061a4fdec9a422d36e0e60a7) docs: update README.md
* [`263f699`](https://github.com/siderolabs/extensions/commit/263f699ea46177f18fd9c9510effad4e663dcb37) chore: annotate extensions with tiers
* [`59496eb`](https://github.com/siderolabs/extensions/commit/59496eb8761a360ca27ea3664277a0b2138a8601) feat: update to Go 1.25 and Linux 6.15
* [`40c7a03`](https://github.com/siderolabs/extensions/commit/40c7a03f1e488f4ebd96b61dac52719f3c1c02eb) feat: update NVIDIA LTS to series 580
* [`ea0d1a3`](https://github.com/siderolabs/extensions/commit/ea0d1a32a4bc947be119891e7cc97c7953ada3e4) feat: update Linux firmware and Intel u-code
* [`1b1d883`](https://github.com/siderolabs/extensions/commit/1b1d883e37eb999e4ed149e9e0252b6dafe1a80c) chore: update nfsrahead defaults
* [`ba3c2ac`](https://github.com/siderolabs/extensions/commit/ba3c2ac3521f3aa9ad667d390021194e7031042a) feat: update Go to 1.24.6
* [`7ef078a`](https://github.com/siderolabs/extensions/commit/7ef078a6324dd2cd9fd1cc319e8cfee84ebf7859) fix: remove sbom of kernel modules
* [`efcb066`](https://github.com/siderolabs/extensions/commit/efcb066a6634dd0bfaa0c45b99eaedc38e1b2d14) docs: add SBOM for more extensions
* [`b42fa41`](https://github.com/siderolabs/extensions/commit/b42fa4142f13a9a06de1b94d238fb295916ff13d) feat: update pkgs for containerd 2.1.4
* [`7f8a3a1`](https://github.com/siderolabs/extensions/commit/7f8a3a18cbaa3f052f816f5d98eb081af92c44a0) feat(ci): add grype scan
* [`35b176d`](https://github.com/siderolabs/extensions/commit/35b176d6e3468718a908b75c64542ca6db0a76b4) fix: nfsrahead udev rule
* [`b4c2a62`](https://github.com/siderolabs/extensions/commit/b4c2a62dc9f5d06b491f39ec9cb6f54202838ebe) docs: add SBOM for more extensions
* [`bbea573`](https://github.com/siderolabs/extensions/commit/bbea57328cfabaa3013ba4deb9f8a5b4634a21b2) feat: add thunderbolt udev rule
* [`c66e678`](https://github.com/siderolabs/extensions/commit/c66e678b2b0868350b53d910cd0d9a9a3458b42c) docs: add SBOM for container-runtimes
</p>
</details>

### Changes since v1.12.0-alpha.1
<details><summary>11 commits</summary>
<p>

* [`fc85294`](https://github.com/siderolabs/extensions/commit/fc85294ac92a7fb3ee6cf99b737e739c63696afb) feat: bump deps
* [`ed8c228`](https://github.com/siderolabs/extensions/commit/ed8c228c4fea1b4515dd5d286798cfcfec3bdb47) feat: update cloudflared to 2025.10.0
* [`0beff58`](https://github.com/siderolabs/extensions/commit/0beff582fc439e159a83a11b76ddef74dc344b4a) feat: update dependencies
* [`9d782c5`](https://github.com/siderolabs/extensions/commit/9d782c572a2fcdc086fecd5aded326772ab78b40) feat: fabricmanager support Blackwell baseboards (DGX/HGX B100/B200/B300)
* [`823659b`](https://github.com/siderolabs/extensions/commit/823659bcaea94a55c54a912b84075777070d50f1) feat: use `image-signer`
* [`3f0e872`](https://github.com/siderolabs/extensions/commit/3f0e8727f353b60ffe8ad651a3804c25ad12e23f) feat: add xe extension
* [`e38736c`](https://github.com/siderolabs/extensions/commit/e38736c2ba56efde6d862e39f6b0b786ec5ba5fa) chore: update external version references
* [`38f9757`](https://github.com/siderolabs/extensions/commit/38f975746c903d10df35ce074c8713deb4e77345) chore: update dependencies
* [`83ee3b0`](https://github.com/siderolabs/extensions/commit/83ee3b07c140e8fe2020a293b5cf4e4c297476ce) feat: update tools and pkgs for new Go and NVIDIA LTS
* [`563da22`](https://github.com/siderolabs/extensions/commit/563da22175e2ba46aab11b9ef47318f9c1caa655) feat: update dependencies to match pkgs
* [`5133393`](https://github.com/siderolabs/extensions/commit/5133393ba9a2e210cd8ed920e8b121aa4022763f) fix: try fixing the grype scan
</p>
</details>

### Changes from siderolabs/pkgs
<details><summary>47 commits</summary>
<p>

* [`da97c36`](https://github.com/siderolabs/pkgs/commit/da97c368be82a0d8effd8eae257a87e5f2e29079) feat: update linux-firmware
* [`6d58d7f`](https://github.com/siderolabs/pkgs/commit/6d58d7f86e6499ac5c229d93e0c370ea6370c5e6) feat: bump deps
* [`b535af8`](https://github.com/siderolabs/pkgs/commit/b535af8b28844dfb88904db531cffc5ec86aa490) feat: update dependencies
* [`a098092`](https://github.com/siderolabs/pkgs/commit/a0980927d55c05aaab3bf3d8439d29304ad88972) feat: update Linux to 6.17.3, tt-kmd to 2.4.1
* [`661e578`](https://github.com/siderolabs/pkgs/commit/661e5788713e20df7be51ae90f76f94e45039bf8) feat: add xe extension
* [`8ddac2d`](https://github.com/siderolabs/pkgs/commit/8ddac2da15d625144cfee4613f04bcdec2224297) feat: bump go
* [`332303e`](https://github.com/siderolabs/pkgs/commit/332303ee2228694fe12979aa3fe03748d4eb4bd0) fix: rollback libseccomp version
* [`f62ebca`](https://github.com/siderolabs/pkgs/commit/f62ebca758af069f23c46817e248aaf1674f39be) chore: update dependencies
* [`56f8ae3`](https://github.com/siderolabs/pkgs/commit/56f8ae370d784ef61420bdd053eceefeaa12e194) feat: update Linux to 6.17.1, NVIDIA LTS to 580.95.05
* [`20b1849`](https://github.com/siderolabs/pkgs/commit/20b1849146c50bce39582455717776d330973579) fix: revert "feat" support adding extra trusted certificates in the kernel"
* [`1e3d375`](https://github.com/siderolabs/pkgs/commit/1e3d37514394b4ed9c7b86b1a55b7ff633252142) feat: bump go
* [`ddfd7af`](https://github.com/siderolabs/pkgs/commit/ddfd7afebe9a73c1abe291c03ec3b2e0cb41c0c5) feat: bump dependencies
* [`4dc7709`](https://github.com/siderolabs/pkgs/commit/4dc770905a581c471c7b6eb0eca85177f5cbdd24) feat: update runc to 1.3.2
* [`61d8b44`](https://github.com/siderolabs/pkgs/commit/61d8b44a153211521c1787ebd2478d05232d9a97) chore: fix renovate config for urcu & hailort
* [`5bda512`](https://github.com/siderolabs/pkgs/commit/5bda51205cee2090e83c33697812d2b80a3c6b67) feat: upgrade Linux to 6.17
* [`202a8e6`](https://github.com/siderolabs/pkgs/commit/202a8e663efaf24b662f349c82c5b6addeb6b3a2) feat: update Linux to 6.16.9
* [`3a0900f`](https://github.com/siderolabs/pkgs/commit/3a0900f8b464f22773a153c65d85c6dd075cadba) feat: enable SRv6 LWTUNNEL and BPF support
* [`628efc8`](https://github.com/siderolabs/pkgs/commit/628efc86c2020464c4ee58f419b8b17c8f87aa45) chore: update linuxfirmware and rekres
* [`9d1fb02`](https://github.com/siderolabs/pkgs/commit/9d1fb029db9837af5529f7de371cc9bc709e8cbe) feat: support adding extra trusted certificates in the kernel
* [`7fe686d`](https://github.com/siderolabs/pkgs/commit/7fe686dec6959a32a6a6d6610cedd18d963412ad) fix: build nftables with embedded gmp
* [`fede0a7`](https://github.com/siderolabs/pkgs/commit/fede0a7ec672e835372b8915896fd0e775da0297) feat: add nft binary
* [`0dae01a`](https://github.com/siderolabs/pkgs/commit/0dae01a735bae635ab987febac42c382d851d7d0) feat: update NVIDIA to 580.82.07
* [`9ac2392`](https://github.com/siderolabs/pkgs/commit/9ac23925cf350075a4931a54176fdd3d9b9b7cb7) feat: enable Kernel config options for IPVS Maglev hashing scheduler support
* [`3c5315c`](https://github.com/siderolabs/pkgs/commit/3c5315cd2ae931eb76233d26283f861ace932b1d) feat: update dependencies
* [`122fa66`](https://github.com/siderolabs/pkgs/commit/122fa6626ff8c300623caa625bdc6bc72e866494) feat: update Linux to 6.16.6
* [`ab1e866`](https://github.com/siderolabs/pkgs/commit/ab1e86612ac903b811ecd3b76d55e85951190565) feat: update Go to 1.25.1
* [`7d6ef1b`](https://github.com/siderolabs/pkgs/commit/7d6ef1b187956a8cb2c837f339169b6cb7eb9c12) feat: update runc to 1.3.1
* [`e067c20`](https://github.com/siderolabs/pkgs/commit/e067c2015302c5069676727f047b6aef9ce0dc0c) feat: enable USB audio support
* [`c4faa38`](https://github.com/siderolabs/pkgs/commit/c4faa389341f6be5dad5fdd575acf61e9eb170f8) feat: bump dependencies
* [`453cdfc`](https://github.com/siderolabs/pkgs/commit/453cdfc2fdea9945c8e3531ef190971211eaaf13) feat: enable ublk support
* [`9824684`](https://github.com/siderolabs/pkgs/commit/982468471d8118ac4653b704512e5f847a148dd7) fix: enable memcg v1
* [`2447e11`](https://github.com/siderolabs/pkgs/commit/2447e11dcbcb5dc10703515e2185f753b04e20e0) feat: update Linux to 6.16, GCC to 15
* [`2cfb920`](https://github.com/siderolabs/pkgs/commit/2cfb920acd88d63c3d3ced3a5760549aa180208c) feat: update Linux to 6.15.11, update tools, rekres
* [`ab4e975`](https://github.com/siderolabs/pkgs/commit/ab4e9755b0e2dbf38c75db5d2ff7720f511fd50c) feat: update Linux to 6.12.43
* [`cd67e36`](https://github.com/siderolabs/pkgs/commit/cd67e3660fa0a2ad25ca4b8dcd3c1ce9b96b0b72) chore: update kernel config to support max SMP CPUs
* [`e3b2094`](https://github.com/siderolabs/pkgs/commit/e3b209474060f5a67e36c9239a3a066ee8ace2fe) fix: fix build for new NVIDIA drivers
* [`fd5fdfd`](https://github.com/siderolabs/pkgs/commit/fd5fdfde0bdc4dfc1e9990300df46b9af23c0dfd) feat: update Nvidia LTS to 580.65.06 and production to 570.172.08
* [`0edf426`](https://github.com/siderolabs/pkgs/commit/0edf426d758d67f7baaaa42facdc658396f02f9f) fix: backport CVE kernel patches to 6.12
* [`26d8fef`](https://github.com/siderolabs/pkgs/commit/26d8fefe10329e8d1c285014af0bffe1b9a65431) feat: enable Infiniband IRDMA support
* [`16b5fac`](https://github.com/siderolabs/pkgs/commit/16b5facdbb37f2ad0329bf131ded62cc9b1239a9) fix: re-enable CPUSETS_V1 cgroups controller
* [`fd53886`](https://github.com/siderolabs/pkgs/commit/fd53886f4f36e73181b7b1a0718801bf8e2aadb9) feat: update backportable dependencies
* [`d5f7467`](https://github.com/siderolabs/pkgs/commit/d5f746715727ec34fca7a87ab9f1fac2051f0f75) feat: update Go to 1.24.6
* [`0bd019f`](https://github.com/siderolabs/pkgs/commit/0bd019f29031b7461fbe49552b88d0e26861f955) feat: update containerd to 2.1.4
* [`0ba8b5b`](https://github.com/siderolabs/pkgs/commit/0ba8b5b49f3dedcc49f4040a6f5c57329f5c5605) feat: enable F71808E watchdog driver
* [`895a86b`](https://github.com/siderolabs/pkgs/commit/895a86bcdfedfd9ca1a698d8f8aa71e3600a22c2) fix: enable ISCSI IBFT
* [`a76a67c`](https://github.com/siderolabs/pkgs/commit/a76a67c860a5100f41223fea936712760b33a4cd) feat: update Linux to 6.12.40
* [`8b0a561`](https://github.com/siderolabs/pkgs/commit/8b0a56180198d360ea71b2c62669545b867f9a67) feat: enable bootloader control on amd64
</p>
</details>

### Changes from siderolabs/tools
<details><summary>17 commits</summary>
<p>

* [`a08cc1f`](https://github.com/siderolabs/tools/commit/a08cc1ff80c045b3683ae06d27d2ea2abcabfad2) feat: update git to 2.51.1
* [`e62d613`](https://github.com/siderolabs/tools/commit/e62d613d3113857052acd0eda598805ab73e3dcd) feat: bump go
* [`916b464`](https://github.com/siderolabs/tools/commit/916b4646151ec97a26ef3d1701d084820c41479c) fix: add pkgconf for ncurses, fix Renovate configs, bump deps
* [`11f0337`](https://github.com/siderolabs/tools/commit/11f0337f796f3e39b60ae2ade4babdf5fb67534f) feat: update Go
* [`2c56d7a`](https://github.com/siderolabs/tools/commit/2c56d7ae2e68295ce4820b19e9dc96d108b76632) feat: update OpenSSL to 3.5.4
* [`8f27cfa`](https://github.com/siderolabs/tools/commit/8f27cfab6454aea38a70a9500bf8859f4d04cc21) feat: update dependencies
* [`1c1420e`](https://github.com/siderolabs/tools/commit/1c1420e33eea2ae3c6b0d054474d9bfe1b34c88f) feat: add tinfo to ncurses
* [`7c7328b`](https://github.com/siderolabs/tools/commit/7c7328b599b4cbac35ed7140798dd61c372e628a) fix: set regex in renovate config directly
* [`3ab353b`](https://github.com/siderolabs/tools/commit/3ab353bc82f0d0c76d5bcc4022c851ae5e802d77) fix: modify renovate regex on ca_certificates
* [`4f90801`](https://github.com/siderolabs/tools/commit/4f908016475257651dc41a4250c2c843b5373d08) chore: update openssl, curl, libexpat and rekres
* [`c37ac80`](https://github.com/siderolabs/tools/commit/c37ac805a17daa8c0dc26da18b768864b821920b) feat: update Go to 1.25.1
* [`7c659e9`](https://github.com/siderolabs/tools/commit/7c659e92db3884737abda95e643995107aa14010) feat: update to GCC 15
* [`83fd7b7`](https://github.com/siderolabs/tools/commit/83fd7b7be62f2f59abeb24c971699895759ebb88) feat: migrate from pkg-config to pkgconf
* [`edafd5f`](https://github.com/siderolabs/tools/commit/edafd5f395b1fd31650270332a871c830a5fd781) feat: update toolchain for new Go and Linux headers
* [`65789c7`](https://github.com/siderolabs/tools/commit/65789c75ebd4020a444789cfd74d35fefc2497c2) chore: drop unused vars from Pkgfile
* [`52db66e`](https://github.com/siderolabs/tools/commit/52db66e8d9e1d83e4872a504d45fdabed31504f7) chore: drop protobuf-related stuff from tools
* [`e3c3ef2`](https://github.com/siderolabs/tools/commit/e3c3ef2b604fb80143a17879eeec9f30ca7b07dd) feat: update Go to 1.24.6
</p>
</details>

### Dependency Changes

* **github.com/siderolabs/pkgs**   v1.11.0-15-g2ac857a -> v1.12.0-alpha.0-45-gda97c36
* **github.com/siderolabs/tools**  v1.11.0-2-g8556c73 -> v1.12.0-alpha.0-16-ga08cc1f

Previous release can be found at [v1.11.0](https://github.com/siderolabs/extensions/releases/tag/v1.11.0)

## [Talos System Extensions 1.12.0-alpha.1](https://github.com/siderolabs/extensions/releases/tag/v1.12.0-alpha.1) (2025-10-01)

Welcome to the v1.12.0-alpha.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.11/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

NVIDIA LTS: 580.82.07
NVIDIA Production: 570.172.08
crun: 1.24
gvisor: 20250820.0
ecr-credential-provider: 1.34.0
kata-containers: 3.21.0
spin: 0.21.0
stargz-snapshotter: 0.17.0
youki: 0.5.5
lldpd: 1.0.20
nebula: 1.9.6
netbird: 0.58.2
newt: 1.5.0
tailscale: 1.88.3
zerotier: 1.16.0
drbd: 9.2.15
fuse3: 3.17.4
nut-client: 2.8.4
talos-vmtoolsd: 1.4.0


### Contributors

* Andrey Smirnov
* Dmitrii Sharshakov
* Mateusz Urbanek
* Noel Georgi
* Steven Kreitzer
* Jorik Jonker
* Breland Miley
* Fred Heinecke
* Guillaume LEGRAIN
* Hugo Meyronneinc
* Max Makarov
* Robin Elfrink
* Sammy ETUR
* Steve Francis
* veymax
* Łukasz Szczepański

### Changes
<details><summary>32 commits</summary>
<p>

* [`6962487`](https://github.com/siderolabs/extensions/commit/6962487685afd667999d6bfc175b646ff9366aba) feat: update extensions
* [`f106fb4`](https://github.com/siderolabs/extensions/commit/f106fb4ac3d0d3db8055938a60c71e799b892f84) feat: add usb-audio-drivers
* [`b382dba`](https://github.com/siderolabs/extensions/commit/b382dba8c747124b4b3985ef2f3d279af155c83d) feat: add netbird
* [`17be2f7`](https://github.com/siderolabs/extensions/commit/17be2f7b15e9c17e21f7fa2d77d250dbd2a0e68e) chore: update linux-firmware
* [`12654ec`](https://github.com/siderolabs/extensions/commit/12654ec141789239fd3aed9f09dc427b269206ad) feat: update vmtoolsd to v1.4.0
* [`47051c6`](https://github.com/siderolabs/extensions/commit/47051c6d5da0ab9daad21f158c54bfbea8817347) feat: update NVIDIA LTS to 580.82.07
* [`caa10a4`](https://github.com/siderolabs/extensions/commit/caa10a495fec59f2a726bf8df728ead8ec74e186) docs: update zfs readme for existing pools
* [`b86a64b`](https://github.com/siderolabs/extensions/commit/b86a64b230e2b77d8850dcb2393b5095dc018c01) feat: update dependencies
* [`ea43af9`](https://github.com/siderolabs/extensions/commit/ea43af91e0d9c730742821b8f58ffeebe9b003ef) feat: update Go to 1.25.1
* [`ffcd027`](https://github.com/siderolabs/extensions/commit/ffcd027e34114cb90bedb76eb53d71c6989b3852) chore: redo dependencies as SVG
* [`ef73a85`](https://github.com/siderolabs/extensions/commit/ef73a85d368405763143382c6e66dfe78e7cbf55) feat: update dependencies
* [`2273c9b`](https://github.com/siderolabs/extensions/commit/2273c9be9c4522499a476d46e3ab239991cb766e) release(v1.12.0-alpha.0): prepare release
* [`8976af3`](https://github.com/siderolabs/extensions/commit/8976af37b40f38b2135847fb88047ec44bd19203) fix: nvidia-container-toolkit: use SHA512 from a variable for tirpc
* [`5de79a6`](https://github.com/siderolabs/extensions/commit/5de79a6ba36aa306bb6145807169d41b23eb2c8b) feat: update to Linux 6.16 and GCC 15
* [`d6bc2ce`](https://github.com/siderolabs/extensions/commit/d6bc2ce44b67a6545e235ab733d20346f11fdc71) docs: fix catalog for qemu-guest-agent
* [`5ccda4b`](https://github.com/siderolabs/extensions/commit/5ccda4bdeb79fb916b52122dcc16df6444984735) feat: zerotier - add possible custom planet file
* [`df7a67d`](https://github.com/siderolabs/extensions/commit/df7a67dc8d985f2b104caf524efc530ea0bc5634) chore: refactor manifest for talos-vmtoolsd
* [`d7a938f`](https://github.com/siderolabs/extensions/commit/d7a938fe06b34247061a4fdec9a422d36e0e60a7) docs: update README.md
* [`263f699`](https://github.com/siderolabs/extensions/commit/263f699ea46177f18fd9c9510effad4e663dcb37) chore: annotate extensions with tiers
* [`59496eb`](https://github.com/siderolabs/extensions/commit/59496eb8761a360ca27ea3664277a0b2138a8601) feat: update to Go 1.25 and Linux 6.15
* [`40c7a03`](https://github.com/siderolabs/extensions/commit/40c7a03f1e488f4ebd96b61dac52719f3c1c02eb) feat: update NVIDIA LTS to series 580
* [`ea0d1a3`](https://github.com/siderolabs/extensions/commit/ea0d1a32a4bc947be119891e7cc97c7953ada3e4) feat: update Linux firmware and Intel u-code
* [`1b1d883`](https://github.com/siderolabs/extensions/commit/1b1d883e37eb999e4ed149e9e0252b6dafe1a80c) chore: update nfsrahead defaults
* [`ba3c2ac`](https://github.com/siderolabs/extensions/commit/ba3c2ac3521f3aa9ad667d390021194e7031042a) feat: update Go to 1.24.6
* [`7ef078a`](https://github.com/siderolabs/extensions/commit/7ef078a6324dd2cd9fd1cc319e8cfee84ebf7859) fix: remove sbom of kernel modules
* [`efcb066`](https://github.com/siderolabs/extensions/commit/efcb066a6634dd0bfaa0c45b99eaedc38e1b2d14) docs: add SBOM for more extensions
* [`b42fa41`](https://github.com/siderolabs/extensions/commit/b42fa4142f13a9a06de1b94d238fb295916ff13d) feat: update pkgs for containerd 2.1.4
* [`7f8a3a1`](https://github.com/siderolabs/extensions/commit/7f8a3a18cbaa3f052f816f5d98eb081af92c44a0) feat(ci): add grype scan
* [`35b176d`](https://github.com/siderolabs/extensions/commit/35b176d6e3468718a908b75c64542ca6db0a76b4) fix: nfsrahead udev rule
* [`b4c2a62`](https://github.com/siderolabs/extensions/commit/b4c2a62dc9f5d06b491f39ec9cb6f54202838ebe) docs: add SBOM for more extensions
* [`bbea573`](https://github.com/siderolabs/extensions/commit/bbea57328cfabaa3013ba4deb9f8a5b4634a21b2) feat: add thunderbolt udev rule
* [`c66e678`](https://github.com/siderolabs/extensions/commit/c66e678b2b0868350b53d910cd0d9a9a3458b42c) docs: add SBOM for container-runtimes
</p>
</details>

### Changes since v1.12.0-alpha.0
<details><summary>11 commits</summary>
<p>

* [`6962487`](https://github.com/siderolabs/extensions/commit/6962487685afd667999d6bfc175b646ff9366aba) feat: update extensions
* [`f106fb4`](https://github.com/siderolabs/extensions/commit/f106fb4ac3d0d3db8055938a60c71e799b892f84) feat: add usb-audio-drivers
* [`b382dba`](https://github.com/siderolabs/extensions/commit/b382dba8c747124b4b3985ef2f3d279af155c83d) feat: add netbird
* [`17be2f7`](https://github.com/siderolabs/extensions/commit/17be2f7b15e9c17e21f7fa2d77d250dbd2a0e68e) chore: update linux-firmware
* [`12654ec`](https://github.com/siderolabs/extensions/commit/12654ec141789239fd3aed9f09dc427b269206ad) feat: update vmtoolsd to v1.4.0
* [`47051c6`](https://github.com/siderolabs/extensions/commit/47051c6d5da0ab9daad21f158c54bfbea8817347) feat: update NVIDIA LTS to 580.82.07
* [`caa10a4`](https://github.com/siderolabs/extensions/commit/caa10a495fec59f2a726bf8df728ead8ec74e186) docs: update zfs readme for existing pools
* [`b86a64b`](https://github.com/siderolabs/extensions/commit/b86a64b230e2b77d8850dcb2393b5095dc018c01) feat: update dependencies
* [`ea43af9`](https://github.com/siderolabs/extensions/commit/ea43af91e0d9c730742821b8f58ffeebe9b003ef) feat: update Go to 1.25.1
* [`ffcd027`](https://github.com/siderolabs/extensions/commit/ffcd027e34114cb90bedb76eb53d71c6989b3852) chore: redo dependencies as SVG
* [`ef73a85`](https://github.com/siderolabs/extensions/commit/ef73a85d368405763143382c6e66dfe78e7cbf55) feat: update dependencies
</p>
</details>

### Changes from siderolabs/pkgs
<details><summary>32 commits</summary>
<p>

* [`202a8e6`](https://github.com/siderolabs/pkgs/commit/202a8e663efaf24b662f349c82c5b6addeb6b3a2) feat: update Linux to 6.16.9
* [`3a0900f`](https://github.com/siderolabs/pkgs/commit/3a0900f8b464f22773a153c65d85c6dd075cadba) feat: enable SRv6 LWTUNNEL and BPF support
* [`628efc8`](https://github.com/siderolabs/pkgs/commit/628efc86c2020464c4ee58f419b8b17c8f87aa45) chore: update linuxfirmware and rekres
* [`9d1fb02`](https://github.com/siderolabs/pkgs/commit/9d1fb029db9837af5529f7de371cc9bc709e8cbe) feat: support adding extra trusted certificates in the kernel
* [`7fe686d`](https://github.com/siderolabs/pkgs/commit/7fe686dec6959a32a6a6d6610cedd18d963412ad) fix: build nftables with embedded gmp
* [`fede0a7`](https://github.com/siderolabs/pkgs/commit/fede0a7ec672e835372b8915896fd0e775da0297) feat: add nft binary
* [`0dae01a`](https://github.com/siderolabs/pkgs/commit/0dae01a735bae635ab987febac42c382d851d7d0) feat: update NVIDIA to 580.82.07
* [`9ac2392`](https://github.com/siderolabs/pkgs/commit/9ac23925cf350075a4931a54176fdd3d9b9b7cb7) feat: enable Kernel config options for IPVS Maglev hashing scheduler support
* [`3c5315c`](https://github.com/siderolabs/pkgs/commit/3c5315cd2ae931eb76233d26283f861ace932b1d) feat: update dependencies
* [`122fa66`](https://github.com/siderolabs/pkgs/commit/122fa6626ff8c300623caa625bdc6bc72e866494) feat: update Linux to 6.16.6
* [`ab1e866`](https://github.com/siderolabs/pkgs/commit/ab1e86612ac903b811ecd3b76d55e85951190565) feat: update Go to 1.25.1
* [`7d6ef1b`](https://github.com/siderolabs/pkgs/commit/7d6ef1b187956a8cb2c837f339169b6cb7eb9c12) feat: update runc to 1.3.1
* [`e067c20`](https://github.com/siderolabs/pkgs/commit/e067c2015302c5069676727f047b6aef9ce0dc0c) feat: enable USB audio support
* [`c4faa38`](https://github.com/siderolabs/pkgs/commit/c4faa389341f6be5dad5fdd575acf61e9eb170f8) feat: bump dependencies
* [`453cdfc`](https://github.com/siderolabs/pkgs/commit/453cdfc2fdea9945c8e3531ef190971211eaaf13) feat: enable ublk support
* [`9824684`](https://github.com/siderolabs/pkgs/commit/982468471d8118ac4653b704512e5f847a148dd7) fix: enable memcg v1
* [`2447e11`](https://github.com/siderolabs/pkgs/commit/2447e11dcbcb5dc10703515e2185f753b04e20e0) feat: update Linux to 6.16, GCC to 15
* [`2cfb920`](https://github.com/siderolabs/pkgs/commit/2cfb920acd88d63c3d3ced3a5760549aa180208c) feat: update Linux to 6.15.11, update tools, rekres
* [`ab4e975`](https://github.com/siderolabs/pkgs/commit/ab4e9755b0e2dbf38c75db5d2ff7720f511fd50c) feat: update Linux to 6.12.43
* [`cd67e36`](https://github.com/siderolabs/pkgs/commit/cd67e3660fa0a2ad25ca4b8dcd3c1ce9b96b0b72) chore: update kernel config to support max SMP CPUs
* [`e3b2094`](https://github.com/siderolabs/pkgs/commit/e3b209474060f5a67e36c9239a3a066ee8ace2fe) fix: fix build for new NVIDIA drivers
* [`fd5fdfd`](https://github.com/siderolabs/pkgs/commit/fd5fdfde0bdc4dfc1e9990300df46b9af23c0dfd) feat: update Nvidia LTS to 580.65.06 and production to 570.172.08
* [`0edf426`](https://github.com/siderolabs/pkgs/commit/0edf426d758d67f7baaaa42facdc658396f02f9f) fix: backport CVE kernel patches to 6.12
* [`26d8fef`](https://github.com/siderolabs/pkgs/commit/26d8fefe10329e8d1c285014af0bffe1b9a65431) feat: enable Infiniband IRDMA support
* [`16b5fac`](https://github.com/siderolabs/pkgs/commit/16b5facdbb37f2ad0329bf131ded62cc9b1239a9) fix: re-enable CPUSETS_V1 cgroups controller
* [`fd53886`](https://github.com/siderolabs/pkgs/commit/fd53886f4f36e73181b7b1a0718801bf8e2aadb9) feat: update backportable dependencies
* [`d5f7467`](https://github.com/siderolabs/pkgs/commit/d5f746715727ec34fca7a87ab9f1fac2051f0f75) feat: update Go to 1.24.6
* [`0bd019f`](https://github.com/siderolabs/pkgs/commit/0bd019f29031b7461fbe49552b88d0e26861f955) feat: update containerd to 2.1.4
* [`0ba8b5b`](https://github.com/siderolabs/pkgs/commit/0ba8b5b49f3dedcc49f4040a6f5c57329f5c5605) feat: enable F71808E watchdog driver
* [`895a86b`](https://github.com/siderolabs/pkgs/commit/895a86bcdfedfd9ca1a698d8f8aa71e3600a22c2) fix: enable ISCSI IBFT
* [`a76a67c`](https://github.com/siderolabs/pkgs/commit/a76a67c860a5100f41223fea936712760b33a4cd) feat: update Linux to 6.12.40
* [`8b0a561`](https://github.com/siderolabs/pkgs/commit/8b0a56180198d360ea71b2c62669545b867f9a67) feat: enable bootloader control on amd64
</p>
</details>

### Changes from siderolabs/tools
<details><summary>8 commits</summary>
<p>

* [`4f90801`](https://github.com/siderolabs/tools/commit/4f908016475257651dc41a4250c2c843b5373d08) chore: update openssl, curl, libexpat and rekres
* [`c37ac80`](https://github.com/siderolabs/tools/commit/c37ac805a17daa8c0dc26da18b768864b821920b) feat: update Go to 1.25.1
* [`7c659e9`](https://github.com/siderolabs/tools/commit/7c659e92db3884737abda95e643995107aa14010) feat: update to GCC 15
* [`83fd7b7`](https://github.com/siderolabs/tools/commit/83fd7b7be62f2f59abeb24c971699895759ebb88) feat: migrate from pkg-config to pkgconf
* [`edafd5f`](https://github.com/siderolabs/tools/commit/edafd5f395b1fd31650270332a871c830a5fd781) feat: update toolchain for new Go and Linux headers
* [`65789c7`](https://github.com/siderolabs/tools/commit/65789c75ebd4020a444789cfd74d35fefc2497c2) chore: drop unused vars from Pkgfile
* [`52db66e`](https://github.com/siderolabs/tools/commit/52db66e8d9e1d83e4872a504d45fdabed31504f7) chore: drop protobuf-related stuff from tools
* [`e3c3ef2`](https://github.com/siderolabs/tools/commit/e3c3ef2b604fb80143a17879eeec9f30ca7b07dd) feat: update Go to 1.24.6
</p>
</details>

### Dependency Changes

* **github.com/siderolabs/pkgs**   v1.11.0-15-g2ac857a -> v1.12.0-alpha.0-30-g202a8e6
* **github.com/siderolabs/tools**  v1.11.0-2-g8556c73 -> v1.12.0-alpha.0-7-g4f90801

Previous release can be found at [v1.11.0](https://github.com/siderolabs/extensions/releases/tag/v1.11.0)

## [Talos System Extensions 1.12.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.12.0-alpha.0) (2025-09-02)

Welcome to the v1.12.0-alpha.0 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.11/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

NVIDIA LTS: 580.65.06
NVIDIA Production: 570.172.08


### Contributors

* Andrey Smirnov
* Dmitrii Sharshakov
* Mateusz Urbanek
* Steven Kreitzer
* Noel Georgi
* Guillaume LEGRAIN
* Hugo Meyronneinc
* Jorik Jonker
* Steve Francis

### Changes
<details><summary>20 commits</summary>
<p>

* [`8976af3`](https://github.com/siderolabs/extensions/commit/8976af37b40f38b2135847fb88047ec44bd19203) fix: nvidia-container-toolkit: use SHA512 from a variable for tirpc
* [`5de79a6`](https://github.com/siderolabs/extensions/commit/5de79a6ba36aa306bb6145807169d41b23eb2c8b) feat: update to Linux 6.16 and GCC 15
* [`d6bc2ce`](https://github.com/siderolabs/extensions/commit/d6bc2ce44b67a6545e235ab733d20346f11fdc71) docs: fix catalog for qemu-guest-agent
* [`5ccda4b`](https://github.com/siderolabs/extensions/commit/5ccda4bdeb79fb916b52122dcc16df6444984735) feat: zerotier - add possible custom planet file
* [`df7a67d`](https://github.com/siderolabs/extensions/commit/df7a67dc8d985f2b104caf524efc530ea0bc5634) chore: refactor manifest for talos-vmtoolsd
* [`d7a938f`](https://github.com/siderolabs/extensions/commit/d7a938fe06b34247061a4fdec9a422d36e0e60a7) docs: update README.md
* [`263f699`](https://github.com/siderolabs/extensions/commit/263f699ea46177f18fd9c9510effad4e663dcb37) chore: annotate extensions with tiers
* [`59496eb`](https://github.com/siderolabs/extensions/commit/59496eb8761a360ca27ea3664277a0b2138a8601) feat: update to Go 1.25 and Linux 6.15
* [`40c7a03`](https://github.com/siderolabs/extensions/commit/40c7a03f1e488f4ebd96b61dac52719f3c1c02eb) feat: update NVIDIA LTS to series 580
* [`ea0d1a3`](https://github.com/siderolabs/extensions/commit/ea0d1a32a4bc947be119891e7cc97c7953ada3e4) feat: update Linux firmware and Intel u-code
* [`1b1d883`](https://github.com/siderolabs/extensions/commit/1b1d883e37eb999e4ed149e9e0252b6dafe1a80c) chore: update nfsrahead defaults
* [`ba3c2ac`](https://github.com/siderolabs/extensions/commit/ba3c2ac3521f3aa9ad667d390021194e7031042a) feat: update Go to 1.24.6
* [`7ef078a`](https://github.com/siderolabs/extensions/commit/7ef078a6324dd2cd9fd1cc319e8cfee84ebf7859) fix: remove sbom of kernel modules
* [`efcb066`](https://github.com/siderolabs/extensions/commit/efcb066a6634dd0bfaa0c45b99eaedc38e1b2d14) docs: add SBOM for more extensions
* [`b42fa41`](https://github.com/siderolabs/extensions/commit/b42fa4142f13a9a06de1b94d238fb295916ff13d) feat: update pkgs for containerd 2.1.4
* [`7f8a3a1`](https://github.com/siderolabs/extensions/commit/7f8a3a18cbaa3f052f816f5d98eb081af92c44a0) feat(ci): add grype scan
* [`35b176d`](https://github.com/siderolabs/extensions/commit/35b176d6e3468718a908b75c64542ca6db0a76b4) fix: nfsrahead udev rule
* [`b4c2a62`](https://github.com/siderolabs/extensions/commit/b4c2a62dc9f5d06b491f39ec9cb6f54202838ebe) docs: add SBOM for more extensions
* [`bbea573`](https://github.com/siderolabs/extensions/commit/bbea57328cfabaa3013ba4deb9f8a5b4634a21b2) feat: add thunderbolt udev rule
* [`c66e678`](https://github.com/siderolabs/extensions/commit/c66e678b2b0868350b53d910cd0d9a9a3458b42c) docs: add SBOM for container-runtimes
</p>
</details>

### Changes from siderolabs/pkgs
<details><summary>16 commits</summary>
<p>

* [`2447e11`](https://github.com/siderolabs/pkgs/commit/2447e11dcbcb5dc10703515e2185f753b04e20e0) feat: update Linux to 6.16, GCC to 15
* [`2cfb920`](https://github.com/siderolabs/pkgs/commit/2cfb920acd88d63c3d3ced3a5760549aa180208c) feat: update Linux to 6.15.11, update tools, rekres
* [`ab4e975`](https://github.com/siderolabs/pkgs/commit/ab4e9755b0e2dbf38c75db5d2ff7720f511fd50c) feat: update Linux to 6.12.43
* [`cd67e36`](https://github.com/siderolabs/pkgs/commit/cd67e3660fa0a2ad25ca4b8dcd3c1ce9b96b0b72) chore: update kernel config to support max SMP CPUs
* [`e3b2094`](https://github.com/siderolabs/pkgs/commit/e3b209474060f5a67e36c9239a3a066ee8ace2fe) fix: fix build for new NVIDIA drivers
* [`fd5fdfd`](https://github.com/siderolabs/pkgs/commit/fd5fdfde0bdc4dfc1e9990300df46b9af23c0dfd) feat: update Nvidia LTS to 580.65.06 and production to 570.172.08
* [`0edf426`](https://github.com/siderolabs/pkgs/commit/0edf426d758d67f7baaaa42facdc658396f02f9f) fix: backport CVE kernel patches to 6.12
* [`26d8fef`](https://github.com/siderolabs/pkgs/commit/26d8fefe10329e8d1c285014af0bffe1b9a65431) feat: enable Infiniband IRDMA support
* [`16b5fac`](https://github.com/siderolabs/pkgs/commit/16b5facdbb37f2ad0329bf131ded62cc9b1239a9) fix: re-enable CPUSETS_V1 cgroups controller
* [`fd53886`](https://github.com/siderolabs/pkgs/commit/fd53886f4f36e73181b7b1a0718801bf8e2aadb9) feat: update backportable dependencies
* [`d5f7467`](https://github.com/siderolabs/pkgs/commit/d5f746715727ec34fca7a87ab9f1fac2051f0f75) feat: update Go to 1.24.6
* [`0bd019f`](https://github.com/siderolabs/pkgs/commit/0bd019f29031b7461fbe49552b88d0e26861f955) feat: update containerd to 2.1.4
* [`0ba8b5b`](https://github.com/siderolabs/pkgs/commit/0ba8b5b49f3dedcc49f4040a6f5c57329f5c5605) feat: enable F71808E watchdog driver
* [`895a86b`](https://github.com/siderolabs/pkgs/commit/895a86bcdfedfd9ca1a698d8f8aa71e3600a22c2) fix: enable ISCSI IBFT
* [`a76a67c`](https://github.com/siderolabs/pkgs/commit/a76a67c860a5100f41223fea936712760b33a4cd) feat: update Linux to 6.12.40
* [`8b0a561`](https://github.com/siderolabs/pkgs/commit/8b0a56180198d360ea71b2c62669545b867f9a67) feat: enable bootloader control on amd64
</p>
</details>

### Changes from siderolabs/tools
<details><summary>6 commits</summary>
<p>

* [`7c659e9`](https://github.com/siderolabs/tools/commit/7c659e92db3884737abda95e643995107aa14010) feat: update to GCC 15
* [`83fd7b7`](https://github.com/siderolabs/tools/commit/83fd7b7be62f2f59abeb24c971699895759ebb88) feat: migrate from pkg-config to pkgconf
* [`edafd5f`](https://github.com/siderolabs/tools/commit/edafd5f395b1fd31650270332a871c830a5fd781) feat: update toolchain for new Go and Linux headers
* [`65789c7`](https://github.com/siderolabs/tools/commit/65789c75ebd4020a444789cfd74d35fefc2497c2) chore: drop unused vars from Pkgfile
* [`52db66e`](https://github.com/siderolabs/tools/commit/52db66e8d9e1d83e4872a504d45fdabed31504f7) chore: drop protobuf-related stuff from tools
* [`e3c3ef2`](https://github.com/siderolabs/tools/commit/e3c3ef2b604fb80143a17879eeec9f30ca7b07dd) feat: update Go to 1.24.6
</p>
</details>

### Dependency Changes

* **github.com/siderolabs/pkgs**   v1.11.0-15-g2ac857a -> v1.12.0-alpha.0-14-g2447e11
* **github.com/siderolabs/tools**  v1.11.0-2-g8556c73 -> v1.12.0-alpha.0-5-g7c659e9

Previous release can be found at [v1.11.0](https://github.com/siderolabs/extensions/releases/tag/v1.11.0)

## [Talos System Extensions 1.11.0-alpha.3](https://github.com/siderolabs/extensions/releases/tag/v1.11.0-alpha.3) (2025-07-02)

Welcome to the v1.11.0-alpha.3 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.11/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### HailoRT drivers

[HailoRT](https://github.com/hailo-ai/hailort-drivers) driver is now supported as an extension.


### Newt client

[Newt](https://github.com/fosrl/newt) client is now supported as an extension.


### Tenstorrent drivers

[Tennstorrent](https://github.com/tenstorrent/tt-kmd) driver is now supported as an extension.


### Component Updates

ZFS: 2.3.3
DRBD: 9.2.14
qemu-guest-agent: 10.0.2
fuse: 3.17.2
nut: 2.8.3
Tailscale: 1.84.0
Linux firmware: 20250613
metal-agent: 0.1.3
Intel u-code: 20250512
wasmedge: 0.6.0
Kata containers: 3.17.0
NVIDIA LTS: 535.247.01
NVIDIA Production: 570.140.08
NVIDIA Container Toolkit: 1.17.8
vmtoolsd: 1.1.0
newt: 1.2.1


### Youki Container Runtime

[Youki](https://github.com/youki-dev/youki) is now supported as a container runtime.


### Contributors

* Andrey Smirnov
* Noel Georgi
* Marco Mihai Condrache
* 0xBrandon
* Andrew Rynhard
* Camp
* Jean-Francois Roy
* Jorik Jonker
* Justin Garrison
* Michael Robbins
* Utku Ozdemir
* Yehia Amer

### Changes
<details><summary>35 commits</summary>
<p>

* [`5895d8b`](https://github.com/siderolabs/extensions/commit/5895d8b4b5d8d5756acde7369b4e489c25a58c4b) release(v1.11.0-alpha.2): prepare release
* [`50e4ba4`](https://github.com/siderolabs/extensions/commit/50e4ba4e132136d0ed5842113d096152ded1b9ac) feat: add DVBSky S952 firmware driver extension
* [`bc53fc3`](https://github.com/siderolabs/extensions/commit/bc53fc3b0c652c7b59aa00f4636a4da83ac7a0d4) feat: add Newt extension
* [`a094039`](https://github.com/siderolabs/extensions/commit/a094039c9344759cd6b213f746a2a896663c0372) feat: update dependencies via pkgs
* [`1f02973`](https://github.com/siderolabs/extensions/commit/1f02973917eb6544c9e6f25631898582a93226da) fix: nvme hostnqn lookup path
* [`0835e02`](https://github.com/siderolabs/extensions/commit/0835e025d2de83d6e9bd9f31f54c698e5665f553) feat: add hailort extension
* [`5c5403b`](https://github.com/siderolabs/extensions/commit/5c5403bb417e3690e2bf99c046db27d9b81405a8) feat: update vmtoolsd to v1.1.0
* [`0723b38`](https://github.com/siderolabs/extensions/commit/0723b38c9146a12d94ef54bcc554c41a1b9dee0b) feat: update to Go 1.24.4
* [`d13039b`](https://github.com/siderolabs/extensions/commit/d13039b66df25a2df5baf4791e6406e5087c1b14) release(v1.11.0-alpha.1): prepare release
* [`c26f543`](https://github.com/siderolabs/extensions/commit/c26f543bbd3313c58c14664f3da88eee657fe76e) fix: glib version
* [`29c809b`](https://github.com/siderolabs/extensions/commit/29c809b0a1eef9c8584931176079dbe282e3fb56) feat: udpate NVIDIA extensions
* [`76bba6d`](https://github.com/siderolabs/extensions/commit/76bba6d7c3f2ee00544198fb7b99dd50a20e3253) feat: bump dependencies
* [`f9b5bf6`](https://github.com/siderolabs/extensions/commit/f9b5bf630d74c172b05335de792ac95228c5683c) fix: nvidia builds
* [`5de8e28`](https://github.com/siderolabs/extensions/commit/5de8e28e8029ea2bf961b703289582de59eca0f1) fix: revert "fix(glibc): use lib64 for better compatibility"
* [`f161267`](https://github.com/siderolabs/extensions/commit/f1612673e2c3c3a97b155d39205a4ef95d2a286e) feat: update Intel u-code to 20250512
* [`1efc06b`](https://github.com/siderolabs/extensions/commit/1efc06bf0e22817e336ceb16c5328b204653c357) fix: nebula extension spec
* [`d668b07`](https://github.com/siderolabs/extensions/commit/d668b07d64ffc21268dfc1b98b60985d38e08203) fix(glibc): use lib64 for better compatibility
* [`30eb717`](https://github.com/siderolabs/extensions/commit/30eb717bb6ed4e546a1e794414d8e409778880cf) chore: bump metal-agent version to v0.1.3
* [`21b44b1`](https://github.com/siderolabs/extensions/commit/21b44b1559c008ac297201d5f22225fa45cb4834) fix: tailscale static builds
* [`5f1070a`](https://github.com/siderolabs/extensions/commit/5f1070a846d10145cf308ae576af3713e9fa3242) fix: staticallty link ecr-credential-provider
* [`40a0b65`](https://github.com/siderolabs/extensions/commit/40a0b65c86dc1b993ac7a73a6eaf45d8e7c361d8) feat: add tenstorrent extension
* [`3df6924`](https://github.com/siderolabs/extensions/commit/3df6924544fccac275e08cb75724e4eb0b47fbc4) feat: update Linux firmware to 20250509
* [`ac5ad5d`](https://github.com/siderolabs/extensions/commit/ac5ad5d76e1ac731a191a59a34dcedd6c8297b2b) feat: update pkgs with containerd 2.1.0
* [`29ce902`](https://github.com/siderolabs/extensions/commit/29ce902bdc66a295196b3349dd416d74b3bf588f) chore: update pkgs and zstd module
* [`9b311a0`](https://github.com/siderolabs/extensions/commit/9b311a06373a397153d1904ce7d411f25845c1a6) feat(panfrost): add kernel drivers
* [`f632ef9`](https://github.com/siderolabs/extensions/commit/f632ef9687bf352b1ffb051d4072fdebf20a10ed) feat: bump dependencies
* [`643c853`](https://github.com/siderolabs/extensions/commit/643c853bbfac1b735d8861eb00d5550a75514a07) fix(panfrost): add mali firmware
* [`e908223`](https://github.com/siderolabs/extensions/commit/e90822373e0a4a1e854f328352e0919c92798501) feat: update for new pkgs, ZFS 2.3.2
* [`bcc1d6e`](https://github.com/siderolabs/extensions/commit/bcc1d6e0cbfcc9045d525db4917c10041ad122b9) fix: lldpd build
* [`573ed8c`](https://github.com/siderolabs/extensions/commit/573ed8c3812ff648906194b6d4f5f7cb48b90c8b) release(v1.11.0-alpha.0): prepare release
* [`0fa2a42`](https://github.com/siderolabs/extensions/commit/0fa2a428f05907ad285bf6f0ac934216e7a3bf87) chore: prepare for 1.11.0-alpha.0 release
* [`08108ca`](https://github.com/siderolabs/extensions/commit/08108cabd4a7fb8947fba460009de6bdf9ee44e3) feat: add youki container runtime extension
* [`e45c086`](https://github.com/siderolabs/extensions/commit/e45c086757e6c9b39fccc6024ee48783956f0285) fix: build of tirpc
* [`6c7422b`](https://github.com/siderolabs/extensions/commit/6c7422b87573bc3919a75bc4dafe71d6bdf06a62) feat: update pkgs
* [`5bf5413`](https://github.com/siderolabs/extensions/commit/5bf5413437d0d3df0dc368b39a9292793f37d698) fix: add SSL mount to Tailscale
</p>
</details>

### Changes since v1.11.0-alpha.2
<details><summary>0 commit</summary>
<p>

</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.10.0](https://github.com/siderolabs/extensions/releases/tag/v1.10.0)

## [Talos System Extensions 1.11.0-alpha.2](https://github.com/siderolabs/extensions/releases/tag/v1.11.0-alpha.2) (2025-07-01)

Welcome to the v1.11.0-alpha.2 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.11/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### HailoRT drivers

[HailoRT](https://github.com/hailo-ai/hailort-drivers) driver is now supported as an extension.


### Newt client

[Newt](https://github.com/fosrl/newt) client is now supported as an extension.


### Tenstorrent drivers

[Tennstorrent](https://github.com/tenstorrent/tt-kmd) driver is now supported as an extension.


### Component Updates

ZFS: 2.3.3
DRBD: 9.2.14
qemu-guest-agent: 10.0.2
fuse: 3.17.2
nut: 2.8.3
Tailscale: 1.84.0
Linux firmware: 20250613
metal-agent: 0.1.3
Intel u-code: 20250512
wasmedge: 0.6.0
Kata containers: 3.17.0
NVIDIA LTS: 535.247.01
NVIDIA Production: 570.140.08
NVIDIA Container Toolkit: 1.17.8
vmtoolsd: 1.1.0
newt: 1.2.1


### Youki Container Runtime

[Youki](https://github.com/youki-dev/youki) is now supported as a container runtime.


### Contributors

* Andrey Smirnov
* Noel Georgi
* Marco Mihai Condrache
* 0xBrandon
* Andrew Rynhard
* Camp
* Jean-Francois Roy
* Jorik Jonker
* Justin Garrison
* Michael Robbins
* Utku Ozdemir
* Yehia Amer

### Changes
<details><summary>34 commits</summary>
<p>

* [`50e4ba4`](https://github.com/siderolabs/extensions/commit/50e4ba4e132136d0ed5842113d096152ded1b9ac) feat: add DVBSky S952 firmware driver extension
* [`bc53fc3`](https://github.com/siderolabs/extensions/commit/bc53fc3b0c652c7b59aa00f4636a4da83ac7a0d4) feat: add Newt extension
* [`a094039`](https://github.com/siderolabs/extensions/commit/a094039c9344759cd6b213f746a2a896663c0372) feat: update dependencies via pkgs
* [`1f02973`](https://github.com/siderolabs/extensions/commit/1f02973917eb6544c9e6f25631898582a93226da) fix: nvme hostnqn lookup path
* [`0835e02`](https://github.com/siderolabs/extensions/commit/0835e025d2de83d6e9bd9f31f54c698e5665f553) feat: add hailort extension
* [`5c5403b`](https://github.com/siderolabs/extensions/commit/5c5403bb417e3690e2bf99c046db27d9b81405a8) feat: update vmtoolsd to v1.1.0
* [`0723b38`](https://github.com/siderolabs/extensions/commit/0723b38c9146a12d94ef54bcc554c41a1b9dee0b) feat: update to Go 1.24.4
* [`d13039b`](https://github.com/siderolabs/extensions/commit/d13039b66df25a2df5baf4791e6406e5087c1b14) release(v1.11.0-alpha.1): prepare release
* [`c26f543`](https://github.com/siderolabs/extensions/commit/c26f543bbd3313c58c14664f3da88eee657fe76e) fix: glib version
* [`29c809b`](https://github.com/siderolabs/extensions/commit/29c809b0a1eef9c8584931176079dbe282e3fb56) feat: udpate NVIDIA extensions
* [`76bba6d`](https://github.com/siderolabs/extensions/commit/76bba6d7c3f2ee00544198fb7b99dd50a20e3253) feat: bump dependencies
* [`f9b5bf6`](https://github.com/siderolabs/extensions/commit/f9b5bf630d74c172b05335de792ac95228c5683c) fix: nvidia builds
* [`5de8e28`](https://github.com/siderolabs/extensions/commit/5de8e28e8029ea2bf961b703289582de59eca0f1) fix: revert "fix(glibc): use lib64 for better compatibility"
* [`f161267`](https://github.com/siderolabs/extensions/commit/f1612673e2c3c3a97b155d39205a4ef95d2a286e) feat: update Intel u-code to 20250512
* [`1efc06b`](https://github.com/siderolabs/extensions/commit/1efc06bf0e22817e336ceb16c5328b204653c357) fix: nebula extension spec
* [`d668b07`](https://github.com/siderolabs/extensions/commit/d668b07d64ffc21268dfc1b98b60985d38e08203) fix(glibc): use lib64 for better compatibility
* [`30eb717`](https://github.com/siderolabs/extensions/commit/30eb717bb6ed4e546a1e794414d8e409778880cf) chore: bump metal-agent version to v0.1.3
* [`21b44b1`](https://github.com/siderolabs/extensions/commit/21b44b1559c008ac297201d5f22225fa45cb4834) fix: tailscale static builds
* [`5f1070a`](https://github.com/siderolabs/extensions/commit/5f1070a846d10145cf308ae576af3713e9fa3242) fix: staticallty link ecr-credential-provider
* [`40a0b65`](https://github.com/siderolabs/extensions/commit/40a0b65c86dc1b993ac7a73a6eaf45d8e7c361d8) feat: add tenstorrent extension
* [`3df6924`](https://github.com/siderolabs/extensions/commit/3df6924544fccac275e08cb75724e4eb0b47fbc4) feat: update Linux firmware to 20250509
* [`ac5ad5d`](https://github.com/siderolabs/extensions/commit/ac5ad5d76e1ac731a191a59a34dcedd6c8297b2b) feat: update pkgs with containerd 2.1.0
* [`29ce902`](https://github.com/siderolabs/extensions/commit/29ce902bdc66a295196b3349dd416d74b3bf588f) chore: update pkgs and zstd module
* [`9b311a0`](https://github.com/siderolabs/extensions/commit/9b311a06373a397153d1904ce7d411f25845c1a6) feat(panfrost): add kernel drivers
* [`f632ef9`](https://github.com/siderolabs/extensions/commit/f632ef9687bf352b1ffb051d4072fdebf20a10ed) feat: bump dependencies
* [`643c853`](https://github.com/siderolabs/extensions/commit/643c853bbfac1b735d8861eb00d5550a75514a07) fix(panfrost): add mali firmware
* [`e908223`](https://github.com/siderolabs/extensions/commit/e90822373e0a4a1e854f328352e0919c92798501) feat: update for new pkgs, ZFS 2.3.2
* [`bcc1d6e`](https://github.com/siderolabs/extensions/commit/bcc1d6e0cbfcc9045d525db4917c10041ad122b9) fix: lldpd build
* [`573ed8c`](https://github.com/siderolabs/extensions/commit/573ed8c3812ff648906194b6d4f5f7cb48b90c8b) release(v1.11.0-alpha.0): prepare release
* [`0fa2a42`](https://github.com/siderolabs/extensions/commit/0fa2a428f05907ad285bf6f0ac934216e7a3bf87) chore: prepare for 1.11.0-alpha.0 release
* [`08108ca`](https://github.com/siderolabs/extensions/commit/08108cabd4a7fb8947fba460009de6bdf9ee44e3) feat: add youki container runtime extension
* [`e45c086`](https://github.com/siderolabs/extensions/commit/e45c086757e6c9b39fccc6024ee48783956f0285) fix: build of tirpc
* [`6c7422b`](https://github.com/siderolabs/extensions/commit/6c7422b87573bc3919a75bc4dafe71d6bdf06a62) feat: update pkgs
* [`5bf5413`](https://github.com/siderolabs/extensions/commit/5bf5413437d0d3df0dc368b39a9292793f37d698) fix: add SSL mount to Tailscale
</p>
</details>

### Changes since v1.11.0-alpha.1
<details><summary>7 commits</summary>
<p>

* [`50e4ba4`](https://github.com/siderolabs/extensions/commit/50e4ba4e132136d0ed5842113d096152ded1b9ac) feat: add DVBSky S952 firmware driver extension
* [`bc53fc3`](https://github.com/siderolabs/extensions/commit/bc53fc3b0c652c7b59aa00f4636a4da83ac7a0d4) feat: add Newt extension
* [`a094039`](https://github.com/siderolabs/extensions/commit/a094039c9344759cd6b213f746a2a896663c0372) feat: update dependencies via pkgs
* [`1f02973`](https://github.com/siderolabs/extensions/commit/1f02973917eb6544c9e6f25631898582a93226da) fix: nvme hostnqn lookup path
* [`0835e02`](https://github.com/siderolabs/extensions/commit/0835e025d2de83d6e9bd9f31f54c698e5665f553) feat: add hailort extension
* [`5c5403b`](https://github.com/siderolabs/extensions/commit/5c5403bb417e3690e2bf99c046db27d9b81405a8) feat: update vmtoolsd to v1.1.0
* [`0723b38`](https://github.com/siderolabs/extensions/commit/0723b38c9146a12d94ef54bcc554c41a1b9dee0b) feat: update to Go 1.24.4
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.10.0](https://github.com/siderolabs/extensions/releases/tag/v1.10.0)

## [Talos System Extensions 1.11.0-alpha.1](https://github.com/siderolabs/extensions/releases/tag/v1.11.0-alpha.1) (2025-06-05)

Welcome to the v1.11.0-alpha.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.11/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Tenstorrent drivers

[Tennstorrent](https://github.com/tenstorrent/tt-kmd) driver is now supported as an extension.


### Component Updates

ZFS: 2.3.2
qemu-guest-agent: 10.0.2
fuse: 3.17.2
nut: 2.8.3
Tailscale: 1.84.0
Linux firmware: 20250509
metal-agent: 0.1.3
Intel u-code: 20250512
wasmedge: 0.6.0
Kata containers: 3.17.0
NVIDIA LTS: 535.247.01
NVIDIA Production: 570.140.08
NVIDIA Container Toolkit: 1.17.8


### Youki Container Runtime

[Youki](https://github.com/youki-dev/youki) is now supported as a container runtime.


### Contributors

* Andrey Smirnov
* Noel Georgi
* Marco Mihai Condrache
* 0xBrandon
* Andrew Rynhard
* Jean-Francois Roy
* Justin Garrison
* Utku Ozdemir

### Changes
<details><summary>26 commits</summary>
<p>

* [`c26f543`](https://github.com/siderolabs/extensions/commit/c26f543bbd3313c58c14664f3da88eee657fe76e) fix: glib version
* [`29c809b`](https://github.com/siderolabs/extensions/commit/29c809b0a1eef9c8584931176079dbe282e3fb56) feat: udpate NVIDIA extensions
* [`76bba6d`](https://github.com/siderolabs/extensions/commit/76bba6d7c3f2ee00544198fb7b99dd50a20e3253) feat: bump dependencies
* [`f9b5bf6`](https://github.com/siderolabs/extensions/commit/f9b5bf630d74c172b05335de792ac95228c5683c) fix: nvidia builds
* [`5de8e28`](https://github.com/siderolabs/extensions/commit/5de8e28e8029ea2bf961b703289582de59eca0f1) fix: revert "fix(glibc): use lib64 for better compatibility"
* [`f161267`](https://github.com/siderolabs/extensions/commit/f1612673e2c3c3a97b155d39205a4ef95d2a286e) feat: update Intel u-code to 20250512
* [`1efc06b`](https://github.com/siderolabs/extensions/commit/1efc06bf0e22817e336ceb16c5328b204653c357) fix: nebula extension spec
* [`d668b07`](https://github.com/siderolabs/extensions/commit/d668b07d64ffc21268dfc1b98b60985d38e08203) fix(glibc): use lib64 for better compatibility
* [`30eb717`](https://github.com/siderolabs/extensions/commit/30eb717bb6ed4e546a1e794414d8e409778880cf) chore: bump metal-agent version to v0.1.3
* [`21b44b1`](https://github.com/siderolabs/extensions/commit/21b44b1559c008ac297201d5f22225fa45cb4834) fix: tailscale static builds
* [`5f1070a`](https://github.com/siderolabs/extensions/commit/5f1070a846d10145cf308ae576af3713e9fa3242) fix: staticallty link ecr-credential-provider
* [`40a0b65`](https://github.com/siderolabs/extensions/commit/40a0b65c86dc1b993ac7a73a6eaf45d8e7c361d8) feat: add tenstorrent extension
* [`3df6924`](https://github.com/siderolabs/extensions/commit/3df6924544fccac275e08cb75724e4eb0b47fbc4) feat: update Linux firmware to 20250509
* [`ac5ad5d`](https://github.com/siderolabs/extensions/commit/ac5ad5d76e1ac731a191a59a34dcedd6c8297b2b) feat: update pkgs with containerd 2.1.0
* [`29ce902`](https://github.com/siderolabs/extensions/commit/29ce902bdc66a295196b3349dd416d74b3bf588f) chore: update pkgs and zstd module
* [`9b311a0`](https://github.com/siderolabs/extensions/commit/9b311a06373a397153d1904ce7d411f25845c1a6) feat(panfrost): add kernel drivers
* [`f632ef9`](https://github.com/siderolabs/extensions/commit/f632ef9687bf352b1ffb051d4072fdebf20a10ed) feat: bump dependencies
* [`643c853`](https://github.com/siderolabs/extensions/commit/643c853bbfac1b735d8861eb00d5550a75514a07) fix(panfrost): add mali firmware
* [`e908223`](https://github.com/siderolabs/extensions/commit/e90822373e0a4a1e854f328352e0919c92798501) feat: update for new pkgs, ZFS 2.3.2
* [`bcc1d6e`](https://github.com/siderolabs/extensions/commit/bcc1d6e0cbfcc9045d525db4917c10041ad122b9) fix: lldpd build
* [`573ed8c`](https://github.com/siderolabs/extensions/commit/573ed8c3812ff648906194b6d4f5f7cb48b90c8b) release(v1.11.0-alpha.0): prepare release
* [`0fa2a42`](https://github.com/siderolabs/extensions/commit/0fa2a428f05907ad285bf6f0ac934216e7a3bf87) chore: prepare for 1.11.0-alpha.0 release
* [`08108ca`](https://github.com/siderolabs/extensions/commit/08108cabd4a7fb8947fba460009de6bdf9ee44e3) feat: add youki container runtime extension
* [`e45c086`](https://github.com/siderolabs/extensions/commit/e45c086757e6c9b39fccc6024ee48783956f0285) fix: build of tirpc
* [`6c7422b`](https://github.com/siderolabs/extensions/commit/6c7422b87573bc3919a75bc4dafe71d6bdf06a62) feat: update pkgs
* [`5bf5413`](https://github.com/siderolabs/extensions/commit/5bf5413437d0d3df0dc368b39a9292793f37d698) fix: add SSL mount to Tailscale
</p>
</details>

### Changes since v1.11.0-alpha.0
<details><summary>20 commits</summary>
<p>

* [`c26f543`](https://github.com/siderolabs/extensions/commit/c26f543bbd3313c58c14664f3da88eee657fe76e) fix: glib version
* [`29c809b`](https://github.com/siderolabs/extensions/commit/29c809b0a1eef9c8584931176079dbe282e3fb56) feat: udpate NVIDIA extensions
* [`76bba6d`](https://github.com/siderolabs/extensions/commit/76bba6d7c3f2ee00544198fb7b99dd50a20e3253) feat: bump dependencies
* [`f9b5bf6`](https://github.com/siderolabs/extensions/commit/f9b5bf630d74c172b05335de792ac95228c5683c) fix: nvidia builds
* [`5de8e28`](https://github.com/siderolabs/extensions/commit/5de8e28e8029ea2bf961b703289582de59eca0f1) fix: revert "fix(glibc): use lib64 for better compatibility"
* [`f161267`](https://github.com/siderolabs/extensions/commit/f1612673e2c3c3a97b155d39205a4ef95d2a286e) feat: update Intel u-code to 20250512
* [`1efc06b`](https://github.com/siderolabs/extensions/commit/1efc06bf0e22817e336ceb16c5328b204653c357) fix: nebula extension spec
* [`d668b07`](https://github.com/siderolabs/extensions/commit/d668b07d64ffc21268dfc1b98b60985d38e08203) fix(glibc): use lib64 for better compatibility
* [`30eb717`](https://github.com/siderolabs/extensions/commit/30eb717bb6ed4e546a1e794414d8e409778880cf) chore: bump metal-agent version to v0.1.3
* [`21b44b1`](https://github.com/siderolabs/extensions/commit/21b44b1559c008ac297201d5f22225fa45cb4834) fix: tailscale static builds
* [`5f1070a`](https://github.com/siderolabs/extensions/commit/5f1070a846d10145cf308ae576af3713e9fa3242) fix: staticallty link ecr-credential-provider
* [`40a0b65`](https://github.com/siderolabs/extensions/commit/40a0b65c86dc1b993ac7a73a6eaf45d8e7c361d8) feat: add tenstorrent extension
* [`3df6924`](https://github.com/siderolabs/extensions/commit/3df6924544fccac275e08cb75724e4eb0b47fbc4) feat: update Linux firmware to 20250509
* [`ac5ad5d`](https://github.com/siderolabs/extensions/commit/ac5ad5d76e1ac731a191a59a34dcedd6c8297b2b) feat: update pkgs with containerd 2.1.0
* [`29ce902`](https://github.com/siderolabs/extensions/commit/29ce902bdc66a295196b3349dd416d74b3bf588f) chore: update pkgs and zstd module
* [`9b311a0`](https://github.com/siderolabs/extensions/commit/9b311a06373a397153d1904ce7d411f25845c1a6) feat(panfrost): add kernel drivers
* [`f632ef9`](https://github.com/siderolabs/extensions/commit/f632ef9687bf352b1ffb051d4072fdebf20a10ed) feat: bump dependencies
* [`643c853`](https://github.com/siderolabs/extensions/commit/643c853bbfac1b735d8861eb00d5550a75514a07) fix(panfrost): add mali firmware
* [`e908223`](https://github.com/siderolabs/extensions/commit/e90822373e0a4a1e854f328352e0919c92798501) feat: update for new pkgs, ZFS 2.3.2
* [`bcc1d6e`](https://github.com/siderolabs/extensions/commit/bcc1d6e0cbfcc9045d525db4917c10041ad122b9) fix: lldpd build
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.10.0](https://github.com/siderolabs/extensions/releases/tag/v1.10.0)

## [Talos System Extensions 1.11.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.11.0-alpha.0) (2025-05-01)

Welcome to the v1.11.0-alpha.0 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.11/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates



### Youki Container Runtime

[Youki](https://github.com/youki-dev/youki) is now supported as a container runtime.


### Contributors

* Andrey Smirnov
* 0xBrandon
* Andrew Rynhard

### Changes
<details><summary>5 commits</summary>
<p>

* [`0fa2a42`](https://github.com/siderolabs/extensions/commit/0fa2a428f05907ad285bf6f0ac934216e7a3bf87) chore: prepare for 1.11.0-alpha.0 release
* [`08108ca`](https://github.com/siderolabs/extensions/commit/08108cabd4a7fb8947fba460009de6bdf9ee44e3) feat: add youki container runtime extension
* [`e45c086`](https://github.com/siderolabs/extensions/commit/e45c086757e6c9b39fccc6024ee48783956f0285) fix: build of tirpc
* [`6c7422b`](https://github.com/siderolabs/extensions/commit/6c7422b87573bc3919a75bc4dafe71d6bdf06a62) feat: update pkgs
* [`5bf5413`](https://github.com/siderolabs/extensions/commit/5bf5413437d0d3df0dc368b39a9292793f37d698) fix: add SSL mount to Tailscale
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.10.0](https://github.com/siderolabs/extensions/releases/tag/v1.10.0)

## [Talos System Extensions 1.10.0-alpha.3](https://github.com/siderolabs/extensions/releases/tag/v1.10.0-alpha.3) (2025-03-24)

Welcome to the v1.10.0-alpha.3 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.10/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Nebula

[Nebula](https://github.com/slackhq/nebula) is now supported as a system extension.


### NFS Server

The NFS server (nfds) as a kernel module is now supported as a system extension.


### Component Updates

Linux Firmware: 20250311
stargz-snapshotter: 0.16.3
ecr-credential-provider: 1.32.1
wasm-edge: 0.5.0
spin: 0.18.0
qemu-guest-agent: 9.2.2
Tailscale: 1.80.3
ZFS: 2.3.1
NVIDIA LTS: 535.230.02
NVIDIA Production: 550.144.03
NVIDIA Container Toolkit: 1.17.4
Intel u-code: 20250211
Gvisor: 20250319.0
Kata Containers: 3.15.0
crun: 1.20
Glibc: 2.41
libseccomp: 2.6.0
Open-iSCSI: 2.1.11


### Contributors

* Andrey Smirnov
* Noel Georgi
* Dmitry Sharshakov
* Utku Ozdemir
* Dmitrii Sharshakov
* Maxime NARBAUD
* Skyler Mäntysaari
* Adam Cirillo
* David Donchez
* Devin Buhl
* Jorik Jonker
* Joseph Glanville
* Lin Yang
* Martin Schuessler
* Rob
* Tom Zaspel
* Will Glynn
* Yevhen Kolomeiko
* iamawacko
* inceabdullah
* mehlc
* 千橘 雫霞

### Changes
<details><summary>53 commits</summary>
<p>

* [`cd1af7f`](https://github.com/siderolabs/extensions/commit/cd1af7f910854bc173ccd91c09a0282a2045eaa0) docs: update README.md
* [`4ef5008`](https://github.com/siderolabs/extensions/commit/4ef5008629335fdcb505ed6166bb52e810214902) feat: update dependencies
* [`aefa9a1`](https://github.com/siderolabs/extensions/commit/aefa9a18dda315157feb1fcc77b6ce9d1c7d6dd8) fix: kata containers build
* [`4bf8841`](https://github.com/siderolabs/extensions/commit/4bf88414a36267c221e1837354d32c98ff9e525e) feat: add helper for supporting the RevolutionPi4
* [`1e2504c`](https://github.com/siderolabs/extensions/commit/1e2504c0935e224b1bb86764e04c1b36b7e73d71) feat: add zerotier extension
* [`69abda0`](https://github.com/siderolabs/extensions/commit/69abda09b5b6f10ca022b266c24deb4f158cb0e6) feat: add extenstion for panfrost module
* [`da519cf`](https://github.com/siderolabs/extensions/commit/da519cf25ad1ef7d882d666ea24ce413293d7f3d) feat: update various extensions
* [`003abea`](https://github.com/siderolabs/extensions/commit/003abea77b99eea5f360c294008e12007bebdeb5) chore: fix Nebula image link in README
* [`d02d972`](https://github.com/siderolabs/extensions/commit/d02d97243155b1cb632580be43d3f7d68fd9864c) chore: unify buildkits
* [`7b3627c`](https://github.com/siderolabs/extensions/commit/7b3627c89b8c93744ca810cde9c5f1381f8584fd) feat: update dependencies
* [`b28f3e1`](https://github.com/siderolabs/extensions/commit/b28f3e1acdc00d751256dc880fb983eb57ee8876) feat: add nfsd extension
* [`edc675e`](https://github.com/siderolabs/extensions/commit/edc675ed726f2feb863645a65134b8c96251d696) release(v1.10.0-alpha.2): prepare release
* [`390ed7f`](https://github.com/siderolabs/extensions/commit/390ed7f8b523613a792367ce9421641a598fcba7) feat: sync kernel version for Talos v1.10.0-alpha.2
* [`e974d6e`](https://github.com/siderolabs/extensions/commit/e974d6e1fcd28b064c37c4c2bec6c3cc8159f6c2) feat: update qlogic firmware for qla2xxx
* [`745e012`](https://github.com/siderolabs/extensions/commit/745e012937fe3e90f8725b9685fb164de3a74969) feat: use tools image from `tools`
* [`8732ec1`](https://github.com/siderolabs/extensions/commit/8732ec1ccfa7631f9500c0b8264ef9648823a776) feat: use reproducibly built Python packages
* [`c6b6c61`](https://github.com/siderolabs/extensions/commit/c6b6c615098ae7252ad2feef1fee67796f9831be) feat: add Nebula extension to Talos
* [`f5c4efa`](https://github.com/siderolabs/extensions/commit/f5c4efaefc3e15dbe7e864e5a1b9bb20dcb25771) docs: correct typo in lldpd extension
* [`bf4027a`](https://github.com/siderolabs/extensions/commit/bf4027a724747f742dd09109fbf33712b9e6557b) chore: remove duplicate commands
* [`c95a427`](https://github.com/siderolabs/extensions/commit/c95a427e54125ec89cc69594dfc93f50a6850fa9) chore: update pkgs
* [`aa141a6`](https://github.com/siderolabs/extensions/commit/aa141a6a7539c99c3f3f3c7116644167033fa1f7) fix: use proper stub for vmtoolsd on arm64
* [`8591d3c`](https://github.com/siderolabs/extensions/commit/8591d3c8ba44435ed81568b92e3ed69717b304bd) fix: update NVIDIA runtime and make its CLI build with current Go
* [`4e31964`](https://github.com/siderolabs/extensions/commit/4e3196407529985a8dd0f956df365181adc416d6) feat: update Intel u-code to 20250211
* [`fd5b270`](https://github.com/siderolabs/extensions/commit/fd5b2700fea4262f98601c9f16d94cac957b50bb) feat: update talos-vmtoolsd to v1.0.0
* [`e12c495`](https://github.com/siderolabs/extensions/commit/e12c4959281b94154c944d42643e6a70d093cc8e) chore: update Linux to 6.12.13
* [`8a17f71`](https://github.com/siderolabs/extensions/commit/8a17f71183c1b6f5347f7fdadfee077209ede48e) feat: bump metal agent to v0.1.2
* [`5cd226e`](https://github.com/siderolabs/extensions/commit/5cd226e3eaef65e9ec8cc902a1ec4f3d4b26f8b1) chore: build with new toolchain
* [`95ddb77`](https://github.com/siderolabs/extensions/commit/95ddb770e659786bf1b6f7ff4c2de233b6ea7f57) feat: mount host ca certs into metal agent
* [`ad72efd`](https://github.com/siderolabs/extensions/commit/ad72efd0cb6845017d018a9242c9ca5056142edf) release(v1.10.0-alpha.1): prepare release
* [`c21e5c4`](https://github.com/siderolabs/extensions/commit/c21e5c4af46bc0dd303e944e207c44dd759773e6) chore: bump pkgs to the latest
* [`70da875`](https://github.com/siderolabs/extensions/commit/70da87510b61ad1a29f22cbdb19dcc30d36e1f91) feat: update dependencies
* [`b1a7dd6`](https://github.com/siderolabs/extensions/commit/b1a7dd670e9b7455b956b98811795c532151e4ec) feat: add nvme cli
* [`4a2e536`](https://github.com/siderolabs/extensions/commit/4a2e5368d4e7db340d2fcb81e0ba38c871489f3a) feat: rework iscsi-tools
* [`ec5e273`](https://github.com/siderolabs/extensions/commit/ec5e273f1aca1dc43ea20c603b82fc904e4bc207) fix: add both Aenix and Enix in MAINTAINERS file
* [`1d51d1d`](https://github.com/siderolabs/extensions/commit/1d51d1d8a4d757d74ed6050d8e323ce251e5dddf) fix: dvb extension as it was missing modules and dep on v4l-uvc-drivers
* [`9c92bda`](https://github.com/siderolabs/extensions/commit/9c92bda95b12c7c00a585778423c8ec8129465c5) fix: unable to override runtime defaults
* [`db466d1`](https://github.com/siderolabs/extensions/commit/db466d1025443ee0b92aa3410b1b20050f64ada9) fix: zfs udev rules installation
* [`e5544b5`](https://github.com/siderolabs/extensions/commit/e5544b5363d5a23c72ecfea2ee3cbff6d7e5dba4) feat: update dependencies
* [`796c40f`](https://github.com/siderolabs/extensions/commit/796c40ff2203cbf1a26b6b12a31c54cf3256dee9) release(v1.10.0-alpha.0): prepare release
* [`1f8bd59`](https://github.com/siderolabs/extensions/commit/1f8bd594049748f125d6d9f8cb52438f58d2eeaf) docs: correct typo
* [`01771bc`](https://github.com/siderolabs/extensions/commit/01771bc305edb2de05c83e80bae1bf5acccd25c6) chore: rekres to simplify `.kres.yaml` defaults
* [`cc467e8`](https://github.com/siderolabs/extensions/commit/cc467e806a5c9eafdf264915eef3eb9bfae2f045) feat: update Linux firmware to 20241210
* [`266346b`](https://github.com/siderolabs/extensions/commit/266346b7142d99670b9f0be2115e462f8490f7cb) chore: rekres for renovate changes
* [`e2b1497`](https://github.com/siderolabs/extensions/commit/e2b149733ebb93023da96551bf286b1e44ac2d0f) fix: nvidia-fabricmanager production
* [`b449434`](https://github.com/siderolabs/extensions/commit/b44943495bc79a1b32a9c232d6c8891a38bab808) fix: add gsc_proxy/mei_gsc_proxy to mei modules
* [`434bd5f`](https://github.com/siderolabs/extensions/commit/434bd5fb78b7bd5b235bd8d3227a5e90bfdbfcd2) fix: use cloudflared release binaries
* [`68c9650`](https://github.com/siderolabs/extensions/commit/68c9650d20da7217986ca91820589236e8c6e64d) chore: bump metal agent version
* [`c7dcaaa`](https://github.com/siderolabs/extensions/commit/c7dcaaa6130a273aac2380a76fee241fb400704c) feat: update ZFS in extensions to 2.2.7
* [`1dd6c36`](https://github.com/siderolabs/extensions/commit/1dd6c364b8aae54a891ef1b371417e4b6b263030) feat: add cloudflared system extension
* [`43efd87`](https://github.com/siderolabs/extensions/commit/43efd87cb106fd9b062ef4ad4f7831eaf7b4fe09) docs: update README.md
* [`ea263ae`](https://github.com/siderolabs/extensions/commit/ea263ae1601e4bf8361119bb68808ed680b2b30d) feat: add dvb-cx23885 extension
* [`4462437`](https://github.com/siderolabs/extensions/commit/44624377575bfecc103dc10a691524a5e2f050d2) docs: replace last command to show the mount
* [`778d80c`](https://github.com/siderolabs/extensions/commit/778d80cd9a9c2be343f5113af722619631334656) feat: set PATH variable in metal-agent
</p>
</details>

### Changes since v1.10.0-alpha.2
<details><summary>11 commits</summary>
<p>

* [`cd1af7f`](https://github.com/siderolabs/extensions/commit/cd1af7f910854bc173ccd91c09a0282a2045eaa0) docs: update README.md
* [`4ef5008`](https://github.com/siderolabs/extensions/commit/4ef5008629335fdcb505ed6166bb52e810214902) feat: update dependencies
* [`aefa9a1`](https://github.com/siderolabs/extensions/commit/aefa9a18dda315157feb1fcc77b6ce9d1c7d6dd8) fix: kata containers build
* [`4bf8841`](https://github.com/siderolabs/extensions/commit/4bf88414a36267c221e1837354d32c98ff9e525e) feat: add helper for supporting the RevolutionPi4
* [`1e2504c`](https://github.com/siderolabs/extensions/commit/1e2504c0935e224b1bb86764e04c1b36b7e73d71) feat: add zerotier extension
* [`69abda0`](https://github.com/siderolabs/extensions/commit/69abda09b5b6f10ca022b266c24deb4f158cb0e6) feat: add extenstion for panfrost module
* [`da519cf`](https://github.com/siderolabs/extensions/commit/da519cf25ad1ef7d882d666ea24ce413293d7f3d) feat: update various extensions
* [`003abea`](https://github.com/siderolabs/extensions/commit/003abea77b99eea5f360c294008e12007bebdeb5) chore: fix Nebula image link in README
* [`d02d972`](https://github.com/siderolabs/extensions/commit/d02d97243155b1cb632580be43d3f7d68fd9864c) chore: unify buildkits
* [`7b3627c`](https://github.com/siderolabs/extensions/commit/7b3627c89b8c93744ca810cde9c5f1381f8584fd) feat: update dependencies
* [`b28f3e1`](https://github.com/siderolabs/extensions/commit/b28f3e1acdc00d751256dc880fb983eb57ee8876) feat: add nfsd extension
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.9.0](https://github.com/siderolabs/extensions/releases/tag/v1.9.0)

## [Talos System Extensions 1.10.0-alpha.2](https://github.com/siderolabs/extensions/releases/tag/v1.10.0-alpha.2) (2025-03-05)

Welcome to the v1.10.0-alpha.2 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.10/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Nebula

[Nebula](https://github.com/slackhq/nebula) is now supported as a system extension.


### Component Updates

Linux Firmware: 20250211
stargz-snapshotter: 0.16.3
ecr-credential-provider: 1.32.0
wasm-edge: 0.5.0
crun: 1.19.1
spin: 0.18.0
qemu-guest-agent: 9.2.0
Tailscale: 1.78.1
ZFS: 2.3.0
NVIDIA LTS: 535.230.02
NVIDIA Production: 550.144.03
NVIDIA Container Toolkit: 1.17.4
Intel u-code: 20250211


### Contributors

* Andrey Smirnov
* Noel Georgi
* Dmitry Sharshakov
* Utku Ozdemir
* Maxime NARBAUD
* Skyler Mäntysaari
* David Donchez
* Devin Buhl
* Dmitrii Sharshakov
* Jorik Jonker
* Tom Zaspel
* Will Glynn
* Yevhen Kolomeiko
* iamawacko
* inceabdullah
* 千橘 雫霞

### Changes
<details><summary>41 commits</summary>
<p>

* [`390ed7f`](https://github.com/siderolabs/extensions/commit/390ed7f8b523613a792367ce9421641a598fcba7) feat: sync kernel version for Talos v1.10.0-alpha.2
* [`e974d6e`](https://github.com/siderolabs/extensions/commit/e974d6e1fcd28b064c37c4c2bec6c3cc8159f6c2) feat: update qlogic firmware for qla2xxx
* [`745e012`](https://github.com/siderolabs/extensions/commit/745e012937fe3e90f8725b9685fb164de3a74969) feat: use tools image from `tools`
* [`8732ec1`](https://github.com/siderolabs/extensions/commit/8732ec1ccfa7631f9500c0b8264ef9648823a776) feat: use reproducibly built Python packages
* [`c6b6c61`](https://github.com/siderolabs/extensions/commit/c6b6c615098ae7252ad2feef1fee67796f9831be) feat: add Nebula extension to Talos
* [`f5c4efa`](https://github.com/siderolabs/extensions/commit/f5c4efaefc3e15dbe7e864e5a1b9bb20dcb25771) docs: correct typo in lldpd extension
* [`bf4027a`](https://github.com/siderolabs/extensions/commit/bf4027a724747f742dd09109fbf33712b9e6557b) chore: remove duplicate commands
* [`c95a427`](https://github.com/siderolabs/extensions/commit/c95a427e54125ec89cc69594dfc93f50a6850fa9) chore: update pkgs
* [`aa141a6`](https://github.com/siderolabs/extensions/commit/aa141a6a7539c99c3f3f3c7116644167033fa1f7) fix: use proper stub for vmtoolsd on arm64
* [`8591d3c`](https://github.com/siderolabs/extensions/commit/8591d3c8ba44435ed81568b92e3ed69717b304bd) fix: update NVIDIA runtime and make its CLI build with current Go
* [`4e31964`](https://github.com/siderolabs/extensions/commit/4e3196407529985a8dd0f956df365181adc416d6) feat: update Intel u-code to 20250211
* [`fd5b270`](https://github.com/siderolabs/extensions/commit/fd5b2700fea4262f98601c9f16d94cac957b50bb) feat: update talos-vmtoolsd to v1.0.0
* [`e12c495`](https://github.com/siderolabs/extensions/commit/e12c4959281b94154c944d42643e6a70d093cc8e) chore: update Linux to 6.12.13
* [`8a17f71`](https://github.com/siderolabs/extensions/commit/8a17f71183c1b6f5347f7fdadfee077209ede48e) feat: bump metal agent to v0.1.2
* [`5cd226e`](https://github.com/siderolabs/extensions/commit/5cd226e3eaef65e9ec8cc902a1ec4f3d4b26f8b1) chore: build with new toolchain
* [`95ddb77`](https://github.com/siderolabs/extensions/commit/95ddb770e659786bf1b6f7ff4c2de233b6ea7f57) feat: mount host ca certs into metal agent
* [`ad72efd`](https://github.com/siderolabs/extensions/commit/ad72efd0cb6845017d018a9242c9ca5056142edf) release(v1.10.0-alpha.1): prepare release
* [`c21e5c4`](https://github.com/siderolabs/extensions/commit/c21e5c4af46bc0dd303e944e207c44dd759773e6) chore: bump pkgs to the latest
* [`70da875`](https://github.com/siderolabs/extensions/commit/70da87510b61ad1a29f22cbdb19dcc30d36e1f91) feat: update dependencies
* [`b1a7dd6`](https://github.com/siderolabs/extensions/commit/b1a7dd670e9b7455b956b98811795c532151e4ec) feat: add nvme cli
* [`4a2e536`](https://github.com/siderolabs/extensions/commit/4a2e5368d4e7db340d2fcb81e0ba38c871489f3a) feat: rework iscsi-tools
* [`ec5e273`](https://github.com/siderolabs/extensions/commit/ec5e273f1aca1dc43ea20c603b82fc904e4bc207) fix: add both Aenix and Enix in MAINTAINERS file
* [`1d51d1d`](https://github.com/siderolabs/extensions/commit/1d51d1d8a4d757d74ed6050d8e323ce251e5dddf) fix: dvb extension as it was missing modules and dep on v4l-uvc-drivers
* [`9c92bda`](https://github.com/siderolabs/extensions/commit/9c92bda95b12c7c00a585778423c8ec8129465c5) fix: unable to override runtime defaults
* [`db466d1`](https://github.com/siderolabs/extensions/commit/db466d1025443ee0b92aa3410b1b20050f64ada9) fix: zfs udev rules installation
* [`e5544b5`](https://github.com/siderolabs/extensions/commit/e5544b5363d5a23c72ecfea2ee3cbff6d7e5dba4) feat: update dependencies
* [`796c40f`](https://github.com/siderolabs/extensions/commit/796c40ff2203cbf1a26b6b12a31c54cf3256dee9) release(v1.10.0-alpha.0): prepare release
* [`1f8bd59`](https://github.com/siderolabs/extensions/commit/1f8bd594049748f125d6d9f8cb52438f58d2eeaf) docs: correct typo
* [`01771bc`](https://github.com/siderolabs/extensions/commit/01771bc305edb2de05c83e80bae1bf5acccd25c6) chore: rekres to simplify `.kres.yaml` defaults
* [`cc467e8`](https://github.com/siderolabs/extensions/commit/cc467e806a5c9eafdf264915eef3eb9bfae2f045) feat: update Linux firmware to 20241210
* [`266346b`](https://github.com/siderolabs/extensions/commit/266346b7142d99670b9f0be2115e462f8490f7cb) chore: rekres for renovate changes
* [`e2b1497`](https://github.com/siderolabs/extensions/commit/e2b149733ebb93023da96551bf286b1e44ac2d0f) fix: nvidia-fabricmanager production
* [`b449434`](https://github.com/siderolabs/extensions/commit/b44943495bc79a1b32a9c232d6c8891a38bab808) fix: add gsc_proxy/mei_gsc_proxy to mei modules
* [`434bd5f`](https://github.com/siderolabs/extensions/commit/434bd5fb78b7bd5b235bd8d3227a5e90bfdbfcd2) fix: use cloudflared release binaries
* [`68c9650`](https://github.com/siderolabs/extensions/commit/68c9650d20da7217986ca91820589236e8c6e64d) chore: bump metal agent version
* [`c7dcaaa`](https://github.com/siderolabs/extensions/commit/c7dcaaa6130a273aac2380a76fee241fb400704c) feat: update ZFS in extensions to 2.2.7
* [`1dd6c36`](https://github.com/siderolabs/extensions/commit/1dd6c364b8aae54a891ef1b371417e4b6b263030) feat: add cloudflared system extension
* [`43efd87`](https://github.com/siderolabs/extensions/commit/43efd87cb106fd9b062ef4ad4f7831eaf7b4fe09) docs: update README.md
* [`ea263ae`](https://github.com/siderolabs/extensions/commit/ea263ae1601e4bf8361119bb68808ed680b2b30d) feat: add dvb-cx23885 extension
* [`4462437`](https://github.com/siderolabs/extensions/commit/44624377575bfecc103dc10a691524a5e2f050d2) docs: replace last command to show the mount
* [`778d80c`](https://github.com/siderolabs/extensions/commit/778d80cd9a9c2be343f5113af722619631334656) feat: set PATH variable in metal-agent
</p>
</details>

### Changes since v1.10.0-alpha.1
<details><summary>16 commits</summary>
<p>

* [`390ed7f`](https://github.com/siderolabs/extensions/commit/390ed7f8b523613a792367ce9421641a598fcba7) feat: sync kernel version for Talos v1.10.0-alpha.2
* [`e974d6e`](https://github.com/siderolabs/extensions/commit/e974d6e1fcd28b064c37c4c2bec6c3cc8159f6c2) feat: update qlogic firmware for qla2xxx
* [`745e012`](https://github.com/siderolabs/extensions/commit/745e012937fe3e90f8725b9685fb164de3a74969) feat: use tools image from `tools`
* [`8732ec1`](https://github.com/siderolabs/extensions/commit/8732ec1ccfa7631f9500c0b8264ef9648823a776) feat: use reproducibly built Python packages
* [`c6b6c61`](https://github.com/siderolabs/extensions/commit/c6b6c615098ae7252ad2feef1fee67796f9831be) feat: add Nebula extension to Talos
* [`f5c4efa`](https://github.com/siderolabs/extensions/commit/f5c4efaefc3e15dbe7e864e5a1b9bb20dcb25771) docs: correct typo in lldpd extension
* [`bf4027a`](https://github.com/siderolabs/extensions/commit/bf4027a724747f742dd09109fbf33712b9e6557b) chore: remove duplicate commands
* [`c95a427`](https://github.com/siderolabs/extensions/commit/c95a427e54125ec89cc69594dfc93f50a6850fa9) chore: update pkgs
* [`aa141a6`](https://github.com/siderolabs/extensions/commit/aa141a6a7539c99c3f3f3c7116644167033fa1f7) fix: use proper stub for vmtoolsd on arm64
* [`8591d3c`](https://github.com/siderolabs/extensions/commit/8591d3c8ba44435ed81568b92e3ed69717b304bd) fix: update NVIDIA runtime and make its CLI build with current Go
* [`4e31964`](https://github.com/siderolabs/extensions/commit/4e3196407529985a8dd0f956df365181adc416d6) feat: update Intel u-code to 20250211
* [`fd5b270`](https://github.com/siderolabs/extensions/commit/fd5b2700fea4262f98601c9f16d94cac957b50bb) feat: update talos-vmtoolsd to v1.0.0
* [`e12c495`](https://github.com/siderolabs/extensions/commit/e12c4959281b94154c944d42643e6a70d093cc8e) chore: update Linux to 6.12.13
* [`8a17f71`](https://github.com/siderolabs/extensions/commit/8a17f71183c1b6f5347f7fdadfee077209ede48e) feat: bump metal agent to v0.1.2
* [`5cd226e`](https://github.com/siderolabs/extensions/commit/5cd226e3eaef65e9ec8cc902a1ec4f3d4b26f8b1) chore: build with new toolchain
* [`95ddb77`](https://github.com/siderolabs/extensions/commit/95ddb770e659786bf1b6f7ff4c2de233b6ea7f57) feat: mount host ca certs into metal agent
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.9.0](https://github.com/siderolabs/extensions/releases/tag/v1.9.0)

## [Talos System Extensions 1.10.0-alpha.1](https://github.com/siderolabs/extensions/releases/tag/v1.10.0-alpha.1) (2025-01-31)

Welcome to the v1.10.0-alpha.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.10/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

Linux Firmware: 20250109
stargz-snapshotter: 0.16.3
ecr-credential-provider: 1.32.0
wasm-edge: 0.5.0
crun: 1.19.1
spin: 0.18.0
qemu-guest-agent: 9.2.0
Tailscale: 1.78.1
ZFS: 2.3.0
NVIDIA LTS: 535.230.02
NVIDIA Production: 550.144.03
NVIDIA Container Toolkit: 1.17.3


### Contributors

* Noel Georgi
* Andrey Smirnov
* Maxime NARBAUD
* Skyler Mäntysaari
* Utku Ozdemir
* David Donchez
* Devin Buhl
* Tom Zaspel
* Will Glynn
* Yevhen Kolomeiko

### Changes
<details><summary>24 commits</summary>
<p>

* [`c21e5c4`](https://github.com/siderolabs/extensions/commit/c21e5c4af46bc0dd303e944e207c44dd759773e6) chore: bump pkgs to the latest
* [`70da875`](https://github.com/siderolabs/extensions/commit/70da87510b61ad1a29f22cbdb19dcc30d36e1f91) feat: update dependencies
* [`b1a7dd6`](https://github.com/siderolabs/extensions/commit/b1a7dd670e9b7455b956b98811795c532151e4ec) feat: add nvme cli
* [`4a2e536`](https://github.com/siderolabs/extensions/commit/4a2e5368d4e7db340d2fcb81e0ba38c871489f3a) feat: rework iscsi-tools
* [`ec5e273`](https://github.com/siderolabs/extensions/commit/ec5e273f1aca1dc43ea20c603b82fc904e4bc207) fix: add both Aenix and Enix in MAINTAINERS file
* [`1d51d1d`](https://github.com/siderolabs/extensions/commit/1d51d1d8a4d757d74ed6050d8e323ce251e5dddf) fix: dvb extension as it was missing modules and dep on v4l-uvc-drivers
* [`9c92bda`](https://github.com/siderolabs/extensions/commit/9c92bda95b12c7c00a585778423c8ec8129465c5) fix: unable to override runtime defaults
* [`db466d1`](https://github.com/siderolabs/extensions/commit/db466d1025443ee0b92aa3410b1b20050f64ada9) fix: zfs udev rules installation
* [`e5544b5`](https://github.com/siderolabs/extensions/commit/e5544b5363d5a23c72ecfea2ee3cbff6d7e5dba4) feat: update dependencies
* [`796c40f`](https://github.com/siderolabs/extensions/commit/796c40ff2203cbf1a26b6b12a31c54cf3256dee9) release(v1.10.0-alpha.0): prepare release
* [`1f8bd59`](https://github.com/siderolabs/extensions/commit/1f8bd594049748f125d6d9f8cb52438f58d2eeaf) docs: correct typo
* [`01771bc`](https://github.com/siderolabs/extensions/commit/01771bc305edb2de05c83e80bae1bf5acccd25c6) chore: rekres to simplify `.kres.yaml` defaults
* [`cc467e8`](https://github.com/siderolabs/extensions/commit/cc467e806a5c9eafdf264915eef3eb9bfae2f045) feat: update Linux firmware to 20241210
* [`266346b`](https://github.com/siderolabs/extensions/commit/266346b7142d99670b9f0be2115e462f8490f7cb) chore: rekres for renovate changes
* [`e2b1497`](https://github.com/siderolabs/extensions/commit/e2b149733ebb93023da96551bf286b1e44ac2d0f) fix: nvidia-fabricmanager production
* [`b449434`](https://github.com/siderolabs/extensions/commit/b44943495bc79a1b32a9c232d6c8891a38bab808) fix: add gsc_proxy/mei_gsc_proxy to mei modules
* [`434bd5f`](https://github.com/siderolabs/extensions/commit/434bd5fb78b7bd5b235bd8d3227a5e90bfdbfcd2) fix: use cloudflared release binaries
* [`68c9650`](https://github.com/siderolabs/extensions/commit/68c9650d20da7217986ca91820589236e8c6e64d) chore: bump metal agent version
* [`c7dcaaa`](https://github.com/siderolabs/extensions/commit/c7dcaaa6130a273aac2380a76fee241fb400704c) feat: update ZFS in extensions to 2.2.7
* [`1dd6c36`](https://github.com/siderolabs/extensions/commit/1dd6c364b8aae54a891ef1b371417e4b6b263030) feat: add cloudflared system extension
* [`43efd87`](https://github.com/siderolabs/extensions/commit/43efd87cb106fd9b062ef4ad4f7831eaf7b4fe09) docs: update README.md
* [`ea263ae`](https://github.com/siderolabs/extensions/commit/ea263ae1601e4bf8361119bb68808ed680b2b30d) feat: add dvb-cx23885 extension
* [`4462437`](https://github.com/siderolabs/extensions/commit/44624377575bfecc103dc10a691524a5e2f050d2) docs: replace last command to show the mount
* [`778d80c`](https://github.com/siderolabs/extensions/commit/778d80cd9a9c2be343f5113af722619631334656) feat: set PATH variable in metal-agent
</p>
</details>

### Changes since v1.10.0-alpha.0
<details><summary>9 commits</summary>
<p>

* [`c21e5c4`](https://github.com/siderolabs/extensions/commit/c21e5c4af46bc0dd303e944e207c44dd759773e6) chore: bump pkgs to the latest
* [`70da875`](https://github.com/siderolabs/extensions/commit/70da87510b61ad1a29f22cbdb19dcc30d36e1f91) feat: update dependencies
* [`b1a7dd6`](https://github.com/siderolabs/extensions/commit/b1a7dd670e9b7455b956b98811795c532151e4ec) feat: add nvme cli
* [`4a2e536`](https://github.com/siderolabs/extensions/commit/4a2e5368d4e7db340d2fcb81e0ba38c871489f3a) feat: rework iscsi-tools
* [`ec5e273`](https://github.com/siderolabs/extensions/commit/ec5e273f1aca1dc43ea20c603b82fc904e4bc207) fix: add both Aenix and Enix in MAINTAINERS file
* [`1d51d1d`](https://github.com/siderolabs/extensions/commit/1d51d1d8a4d757d74ed6050d8e323ce251e5dddf) fix: dvb extension as it was missing modules and dep on v4l-uvc-drivers
* [`9c92bda`](https://github.com/siderolabs/extensions/commit/9c92bda95b12c7c00a585778423c8ec8129465c5) fix: unable to override runtime defaults
* [`db466d1`](https://github.com/siderolabs/extensions/commit/db466d1025443ee0b92aa3410b1b20050f64ada9) fix: zfs udev rules installation
* [`e5544b5`](https://github.com/siderolabs/extensions/commit/e5544b5363d5a23c72ecfea2ee3cbff6d7e5dba4) feat: update dependencies
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.9.0](https://github.com/siderolabs/extensions/releases/tag/v1.9.0)

## [Talos System Extensions 1.10.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.10.0-alpha.0) (2024-12-23)

Welcome to the v1.10.0-alpha.0 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.10/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

Linux Firmware: 20241210


### Contributors

* Noel Georgi
* Andrey Smirnov
* Maxime NARBAUD
* Utku Ozdemir
* Devin Buhl
* Skyler Mäntysaari
* Tom Zaspel
* Will Glynn
* Yevhen Kolomeiko

### Changes
<details><summary>14 commits</summary>
<p>

* [`1f8bd59`](https://github.com/siderolabs/extensions/commit/1f8bd594049748f125d6d9f8cb52438f58d2eeaf) docs: correct typo
* [`01771bc`](https://github.com/siderolabs/extensions/commit/01771bc305edb2de05c83e80bae1bf5acccd25c6) chore: rekres to simplify `.kres.yaml` defaults
* [`cc467e8`](https://github.com/siderolabs/extensions/commit/cc467e806a5c9eafdf264915eef3eb9bfae2f045) feat: update Linux firmware to 20241210
* [`266346b`](https://github.com/siderolabs/extensions/commit/266346b7142d99670b9f0be2115e462f8490f7cb) chore: rekres for renovate changes
* [`e2b1497`](https://github.com/siderolabs/extensions/commit/e2b149733ebb93023da96551bf286b1e44ac2d0f) fix: nvidia-fabricmanager production
* [`b449434`](https://github.com/siderolabs/extensions/commit/b44943495bc79a1b32a9c232d6c8891a38bab808) fix: add gsc_proxy/mei_gsc_proxy to mei modules
* [`434bd5f`](https://github.com/siderolabs/extensions/commit/434bd5fb78b7bd5b235bd8d3227a5e90bfdbfcd2) fix: use cloudflared release binaries
* [`68c9650`](https://github.com/siderolabs/extensions/commit/68c9650d20da7217986ca91820589236e8c6e64d) chore: bump metal agent version
* [`c7dcaaa`](https://github.com/siderolabs/extensions/commit/c7dcaaa6130a273aac2380a76fee241fb400704c) feat: update ZFS in extensions to 2.2.7
* [`1dd6c36`](https://github.com/siderolabs/extensions/commit/1dd6c364b8aae54a891ef1b371417e4b6b263030) feat: add cloudflared system extension
* [`43efd87`](https://github.com/siderolabs/extensions/commit/43efd87cb106fd9b062ef4ad4f7831eaf7b4fe09) docs: update README.md
* [`ea263ae`](https://github.com/siderolabs/extensions/commit/ea263ae1601e4bf8361119bb68808ed680b2b30d) feat: add dvb-cx23885 extension
* [`4462437`](https://github.com/siderolabs/extensions/commit/44624377575bfecc103dc10a691524a5e2f050d2) docs: replace last command to show the mount
* [`778d80c`](https://github.com/siderolabs/extensions/commit/778d80cd9a9c2be343f5113af722619631334656) feat: set PATH variable in metal-agent
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.9.0](https://github.com/siderolabs/extensions/releases/tag/v1.9.0)

## [Talos System Extensions 1.9.0-alpha.3](https://github.com/siderolabs/extensions/releases/tag/v1.9.0-alpha.3) (2024-11-25)

Welcome to the v1.9.0-alpha.3 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.9/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Direct Rendering Manager (DRM)

New `i915` and `amdgpu` extensions are now available.
This combines the previous `i915-ucode` and `amdgpu-firmware` extensions along with the matching kernel modules.
Upgrades via Image Factory will automatically include the new extensions if previously `i915-ucode` or `amdgpu-firmware` were used.


### LLDP

lldpd is now available as a system extension.


### Component Updates

Linux Firmware: 20241110
Tailscale: 1.76.0
crun: 1.17
gvisor: 20241007.0
spin: 0.16.0
ecr-credential-provider: 1.31.1
Intel microcode: 20241112
NVIDIA LTS: 535.216.01
NVIDIA Production: 550.127.05
vmtoolsd-guest-agent: v0.6.1


### Contributors

* Andrey Smirnov
* Noel Georgi
* Jean-Francois Roy
* Tom Zaspel
* Dmitry Sharshakov
* Utku Ozdemir
* Andrei Kvapil
* Mihkel
* Niklas Wik
* Zackeus Bengtsson

### Changes
<details><summary>30 commits</summary>
<p>

* [`f925a59`](https://github.com/siderolabs/extensions/commit/f925a59982c04be9a28a4d0c5ab03502d88bfb8b) chore: combine i915/amdgpu firmware+drivers
* [`9753dd5`](https://github.com/siderolabs/extensions/commit/9753dd59f664947d5994a16c3ce15056cfb028c7) chore: bump extensions validator
* [`8b68b70`](https://github.com/siderolabs/extensions/commit/8b68b704c3396abb9797e90975a964d781904f95) feat: update Intel u-code to 20241112
* [`5f506b6`](https://github.com/siderolabs/extensions/commit/5f506b6ebf9272c3960ff65dd409907ccff12858) feat: update extensions for Linux firmware 20241110
* [`cce8fc6`](https://github.com/siderolabs/extensions/commit/cce8fc6690a6eaa972342de03eed362ce16a897f) chore: bump vmtoolsd version to v0.6.1
* [`3be4b54`](https://github.com/siderolabs/extensions/commit/3be4b54c31608e7339027cf90c828d1638594dd2) release(v1.9.0-alpha.2): prepare release
* [`a32e15d`](https://github.com/siderolabs/extensions/commit/a32e15d80c96646d0ba38fec552f0020f087a76f) release(v1.9.0-alpha.1): prepare release
* [`61c0dc4`](https://github.com/siderolabs/extensions/commit/61c0dc4e5140b8c12a7ec12e8952d27e01404249) feat: add zfs-service to zfs extension (unmount, encryption)
* [`c08262d`](https://github.com/siderolabs/extensions/commit/c08262d8b2b005c3a4174522a956a3a36e3dcadf) feat: update gcc to 14.2, bump NVIDIA driver versions
* [`e3012b8`](https://github.com/siderolabs/extensions/commit/e3012b8e6ca4b49351c5385d74e8e8119f7532fb) chore: add `metal-agent` into readme and reproducibility, rekres
* [`e2a26c0`](https://github.com/siderolabs/extensions/commit/e2a26c0b1cd8b0801be95e810d2fb89eaa0b6777) feat: add extension `metal-agent`
* [`966aaed`](https://github.com/siderolabs/extensions/commit/966aaed644ff7c4837e6c627b437d43ab78c3cc9) feat(fabricmanager): add all topology files
* [`21c0137`](https://github.com/siderolabs/extensions/commit/21c01378aaebbd5c8b599f86007375be6886d4ec) chore: update wolfi-base
* [`7dab4c8`](https://github.com/siderolabs/extensions/commit/7dab4c8751b6bb7289d312b657516c5e873789fb) feat: update Intel u-code to 20241029
* [`6d5f49b`](https://github.com/siderolabs/extensions/commit/6d5f49b92c1a2aed4700ec21e57b2092c6904236) chore: move udev rules
* [`b4311ac`](https://github.com/siderolabs/extensions/commit/b4311accd8750de1d84d380977816d34a23858a1) docs: zfs raid docs
* [`76543d1`](https://github.com/siderolabs/extensions/commit/76543d1cb2813e6d7c64a6012fcc50bd2caab92c) docs: replace nsenter with chroot in README.md
* [`f071725`](https://github.com/siderolabs/extensions/commit/f0717255a9f0c55eff1e214d8b370d99cf8103c0) docs: add steps to create mdadm software raid
* [`544ebcf`](https://github.com/siderolabs/extensions/commit/544ebcfef798f97708cb09d33f6afe2dd8986171) fix: iscsi-tools extension
* [`8e296ab`](https://github.com/siderolabs/extensions/commit/8e296ab29000e6550e7c516cd98c07e6e7f3750e) docs: fix Aenix company name
* [`c7eb377`](https://github.com/siderolabs/extensions/commit/c7eb3771943eba55ec5fe61f7a31e6fce61d73d5) feat: glibc extension
* [`4a66da5`](https://github.com/siderolabs/extensions/commit/4a66da5aa2f28bd71784e5312a6fbc91be4f09d7) release(v1.9.0-alpha.0): prepare release
* [`862d0ac`](https://github.com/siderolabs/extensions/commit/862d0ac63d3a491c87283c9a7423f4d516e42b6f) feat: update dependencies
* [`8a7635b`](https://github.com/siderolabs/extensions/commit/8a7635b2c9ce2cb02fd389192da3d1392ca37674) feat: introduce LLDPD extension service
* [`6a184b8`](https://github.com/siderolabs/extensions/commit/6a184b8acc85237fa0e929277ad028ff576fe4b5) chore: update packages, prepare for v1.9
* [`d474848`](https://github.com/siderolabs/extensions/commit/d4748485fd6d5e44650259fa42bb228fc5cfef12) fix: zfs extensions service yaml to proper path
* [`bfaee18`](https://github.com/siderolabs/extensions/commit/bfaee18773163c35d8a8c91232fb22eaa0732619) feat: update pkgs/Linux firmware
* [`11f48c5`](https://github.com/siderolabs/extensions/commit/11f48c567ff2d491f8dd27897befcdaf87d7c989) fix: image reproducibility with finalize
* [`39d2f20`](https://github.com/siderolabs/extensions/commit/39d2f209d142bf88745aadae74df0830cd865ea4) feat: update Intel u-code to 20240910
* [`8a80e47`](https://github.com/siderolabs/extensions/commit/8a80e47219eb20c4cff2a71995984bdd7fed59fc) fix: reproducibility tests
</p>
</details>

### Changes since v1.9.0-alpha.2
<details><summary>5 commits</summary>
<p>

* [`f925a59`](https://github.com/siderolabs/extensions/commit/f925a59982c04be9a28a4d0c5ab03502d88bfb8b) chore: combine i915/amdgpu firmware+drivers
* [`9753dd5`](https://github.com/siderolabs/extensions/commit/9753dd59f664947d5994a16c3ce15056cfb028c7) chore: bump extensions validator
* [`8b68b70`](https://github.com/siderolabs/extensions/commit/8b68b704c3396abb9797e90975a964d781904f95) feat: update Intel u-code to 20241112
* [`5f506b6`](https://github.com/siderolabs/extensions/commit/5f506b6ebf9272c3960ff65dd409907ccff12858) feat: update extensions for Linux firmware 20241110
* [`cce8fc6`](https://github.com/siderolabs/extensions/commit/cce8fc6690a6eaa972342de03eed362ce16a897f) chore: bump vmtoolsd version to v0.6.1
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.8.0](https://github.com/siderolabs/extensions/releases/tag/v1.8.0)

## [Talos System Extensions 1.9.0-alpha.2](https://github.com/siderolabs/extensions/releases/tag/v1.9.0-alpha.2) (2024-11-08)

Welcome to the v1.9.0-alpha.2 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.9/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### LLDP

lldpd is now available as a system extension.


### Component Updates

Linux Firmware: 20241017
Tailscale: 1.76.0
crun: 1.17
gvisor: 20241007.0
spin: 0.16.0
ecr-credential-provider: 1.31.1
Intel microcode: 20241029
NVIDIA LTS: 535.216.01
NVIDIA Production: 550.127.05


### Contributors

* Andrey Smirnov
* Jean-Francois Roy
* Tom Zaspel
* Dmitry Sharshakov
* Noel Georgi
* Utku Ozdemir
* Andrei Kvapil
* Niklas Wik
* Zackeus Bengtsson

### Changes
<details><summary>24 commits</summary>
<p>

* [`a32e15d`](https://github.com/siderolabs/extensions/commit/a32e15d80c96646d0ba38fec552f0020f087a76f) release(v1.9.0-alpha.1): prepare release
* [`61c0dc4`](https://github.com/siderolabs/extensions/commit/61c0dc4e5140b8c12a7ec12e8952d27e01404249) feat: add zfs-service to zfs extension (unmount, encryption)
* [`c08262d`](https://github.com/siderolabs/extensions/commit/c08262d8b2b005c3a4174522a956a3a36e3dcadf) feat: update gcc to 14.2, bump NVIDIA driver versions
* [`e3012b8`](https://github.com/siderolabs/extensions/commit/e3012b8e6ca4b49351c5385d74e8e8119f7532fb) chore: add `metal-agent` into readme and reproducibility, rekres
* [`e2a26c0`](https://github.com/siderolabs/extensions/commit/e2a26c0b1cd8b0801be95e810d2fb89eaa0b6777) feat: add extension `metal-agent`
* [`966aaed`](https://github.com/siderolabs/extensions/commit/966aaed644ff7c4837e6c627b437d43ab78c3cc9) feat(fabricmanager): add all topology files
* [`21c0137`](https://github.com/siderolabs/extensions/commit/21c01378aaebbd5c8b599f86007375be6886d4ec) chore: update wolfi-base
* [`7dab4c8`](https://github.com/siderolabs/extensions/commit/7dab4c8751b6bb7289d312b657516c5e873789fb) feat: update Intel u-code to 20241029
* [`6d5f49b`](https://github.com/siderolabs/extensions/commit/6d5f49b92c1a2aed4700ec21e57b2092c6904236) chore: move udev rules
* [`b4311ac`](https://github.com/siderolabs/extensions/commit/b4311accd8750de1d84d380977816d34a23858a1) docs: zfs raid docs
* [`76543d1`](https://github.com/siderolabs/extensions/commit/76543d1cb2813e6d7c64a6012fcc50bd2caab92c) docs: replace nsenter with chroot in README.md
* [`f071725`](https://github.com/siderolabs/extensions/commit/f0717255a9f0c55eff1e214d8b370d99cf8103c0) docs: add steps to create mdadm software raid
* [`544ebcf`](https://github.com/siderolabs/extensions/commit/544ebcfef798f97708cb09d33f6afe2dd8986171) fix: iscsi-tools extension
* [`8e296ab`](https://github.com/siderolabs/extensions/commit/8e296ab29000e6550e7c516cd98c07e6e7f3750e) docs: fix Aenix company name
* [`c7eb377`](https://github.com/siderolabs/extensions/commit/c7eb3771943eba55ec5fe61f7a31e6fce61d73d5) feat: glibc extension
* [`4a66da5`](https://github.com/siderolabs/extensions/commit/4a66da5aa2f28bd71784e5312a6fbc91be4f09d7) release(v1.9.0-alpha.0): prepare release
* [`862d0ac`](https://github.com/siderolabs/extensions/commit/862d0ac63d3a491c87283c9a7423f4d516e42b6f) feat: update dependencies
* [`8a7635b`](https://github.com/siderolabs/extensions/commit/8a7635b2c9ce2cb02fd389192da3d1392ca37674) feat: introduce LLDPD extension service
* [`6a184b8`](https://github.com/siderolabs/extensions/commit/6a184b8acc85237fa0e929277ad028ff576fe4b5) chore: update packages, prepare for v1.9
* [`d474848`](https://github.com/siderolabs/extensions/commit/d4748485fd6d5e44650259fa42bb228fc5cfef12) fix: zfs extensions service yaml to proper path
* [`bfaee18`](https://github.com/siderolabs/extensions/commit/bfaee18773163c35d8a8c91232fb22eaa0732619) feat: update pkgs/Linux firmware
* [`11f48c5`](https://github.com/siderolabs/extensions/commit/11f48c567ff2d491f8dd27897befcdaf87d7c989) fix: image reproducibility with finalize
* [`39d2f20`](https://github.com/siderolabs/extensions/commit/39d2f209d142bf88745aadae74df0830cd865ea4) feat: update Intel u-code to 20240910
* [`8a80e47`](https://github.com/siderolabs/extensions/commit/8a80e47219eb20c4cff2a71995984bdd7fed59fc) fix: reproducibility tests
</p>
</details>

### Changes since v1.9.0-alpha.1
<details><summary>0 commit</summary>
<p>

</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.8.0](https://github.com/siderolabs/extensions/releases/tag/v1.8.0)

## [Talos System Extensions 1.9.0-alpha.1](https://github.com/siderolabs/extensions/releases/tag/v1.9.0-alpha.1) (2024-11-08)

Welcome to the v1.9.0-alpha.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.9/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### LLDP

lldpd is now available as a system extension.


### Component Updates

Linux Firmware: 20241017
Tailscale: 1.76.0
crun: 1.17
gvisor: 20241007.0
spin: 0.16.0
ecr-credential-provider: 1.31.1
Intel microcode: 20241029
NVIDIA LTS: 535.216.01
NVIDIA Production: 550.127.05


### Contributors

* Andrey Smirnov
* Jean-Francois Roy
* Tom Zaspel
* Dmitry Sharshakov
* Noel Georgi
* Utku Ozdemir
* Andrei Kvapil
* Niklas Wik
* Zackeus Bengtsson

### Changes
<details><summary>23 commits</summary>
<p>

* [`61c0dc4`](https://github.com/siderolabs/extensions/commit/61c0dc4e5140b8c12a7ec12e8952d27e01404249) feat: add zfs-service to zfs extension (unmount, encryption)
* [`c08262d`](https://github.com/siderolabs/extensions/commit/c08262d8b2b005c3a4174522a956a3a36e3dcadf) feat: update gcc to 14.2, bump NVIDIA driver versions
* [`e3012b8`](https://github.com/siderolabs/extensions/commit/e3012b8e6ca4b49351c5385d74e8e8119f7532fb) chore: add `metal-agent` into readme and reproducibility, rekres
* [`e2a26c0`](https://github.com/siderolabs/extensions/commit/e2a26c0b1cd8b0801be95e810d2fb89eaa0b6777) feat: add extension `metal-agent`
* [`966aaed`](https://github.com/siderolabs/extensions/commit/966aaed644ff7c4837e6c627b437d43ab78c3cc9) feat(fabricmanager): add all topology files
* [`21c0137`](https://github.com/siderolabs/extensions/commit/21c01378aaebbd5c8b599f86007375be6886d4ec) chore: update wolfi-base
* [`7dab4c8`](https://github.com/siderolabs/extensions/commit/7dab4c8751b6bb7289d312b657516c5e873789fb) feat: update Intel u-code to 20241029
* [`6d5f49b`](https://github.com/siderolabs/extensions/commit/6d5f49b92c1a2aed4700ec21e57b2092c6904236) chore: move udev rules
* [`b4311ac`](https://github.com/siderolabs/extensions/commit/b4311accd8750de1d84d380977816d34a23858a1) docs: zfs raid docs
* [`76543d1`](https://github.com/siderolabs/extensions/commit/76543d1cb2813e6d7c64a6012fcc50bd2caab92c) docs: replace nsenter with chroot in README.md
* [`f071725`](https://github.com/siderolabs/extensions/commit/f0717255a9f0c55eff1e214d8b370d99cf8103c0) docs: add steps to create mdadm software raid
* [`544ebcf`](https://github.com/siderolabs/extensions/commit/544ebcfef798f97708cb09d33f6afe2dd8986171) fix: iscsi-tools extension
* [`8e296ab`](https://github.com/siderolabs/extensions/commit/8e296ab29000e6550e7c516cd98c07e6e7f3750e) docs: fix Aenix company name
* [`c7eb377`](https://github.com/siderolabs/extensions/commit/c7eb3771943eba55ec5fe61f7a31e6fce61d73d5) feat: glibc extension
* [`4a66da5`](https://github.com/siderolabs/extensions/commit/4a66da5aa2f28bd71784e5312a6fbc91be4f09d7) release(v1.9.0-alpha.0): prepare release
* [`862d0ac`](https://github.com/siderolabs/extensions/commit/862d0ac63d3a491c87283c9a7423f4d516e42b6f) feat: update dependencies
* [`8a7635b`](https://github.com/siderolabs/extensions/commit/8a7635b2c9ce2cb02fd389192da3d1392ca37674) feat: introduce LLDPD extension service
* [`6a184b8`](https://github.com/siderolabs/extensions/commit/6a184b8acc85237fa0e929277ad028ff576fe4b5) chore: update packages, prepare for v1.9
* [`d474848`](https://github.com/siderolabs/extensions/commit/d4748485fd6d5e44650259fa42bb228fc5cfef12) fix: zfs extensions service yaml to proper path
* [`bfaee18`](https://github.com/siderolabs/extensions/commit/bfaee18773163c35d8a8c91232fb22eaa0732619) feat: update pkgs/Linux firmware
* [`11f48c5`](https://github.com/siderolabs/extensions/commit/11f48c567ff2d491f8dd27897befcdaf87d7c989) fix: image reproducibility with finalize
* [`39d2f20`](https://github.com/siderolabs/extensions/commit/39d2f209d142bf88745aadae74df0830cd865ea4) feat: update Intel u-code to 20240910
* [`8a80e47`](https://github.com/siderolabs/extensions/commit/8a80e47219eb20c4cff2a71995984bdd7fed59fc) fix: reproducibility tests
</p>
</details>

### Changes since v1.9.0-alpha.0
<details><summary>14 commits</summary>
<p>

* [`61c0dc4`](https://github.com/siderolabs/extensions/commit/61c0dc4e5140b8c12a7ec12e8952d27e01404249) feat: add zfs-service to zfs extension (unmount, encryption)
* [`c08262d`](https://github.com/siderolabs/extensions/commit/c08262d8b2b005c3a4174522a956a3a36e3dcadf) feat: update gcc to 14.2, bump NVIDIA driver versions
* [`e3012b8`](https://github.com/siderolabs/extensions/commit/e3012b8e6ca4b49351c5385d74e8e8119f7532fb) chore: add `metal-agent` into readme and reproducibility, rekres
* [`e2a26c0`](https://github.com/siderolabs/extensions/commit/e2a26c0b1cd8b0801be95e810d2fb89eaa0b6777) feat: add extension `metal-agent`
* [`966aaed`](https://github.com/siderolabs/extensions/commit/966aaed644ff7c4837e6c627b437d43ab78c3cc9) feat(fabricmanager): add all topology files
* [`21c0137`](https://github.com/siderolabs/extensions/commit/21c01378aaebbd5c8b599f86007375be6886d4ec) chore: update wolfi-base
* [`7dab4c8`](https://github.com/siderolabs/extensions/commit/7dab4c8751b6bb7289d312b657516c5e873789fb) feat: update Intel u-code to 20241029
* [`6d5f49b`](https://github.com/siderolabs/extensions/commit/6d5f49b92c1a2aed4700ec21e57b2092c6904236) chore: move udev rules
* [`b4311ac`](https://github.com/siderolabs/extensions/commit/b4311accd8750de1d84d380977816d34a23858a1) docs: zfs raid docs
* [`76543d1`](https://github.com/siderolabs/extensions/commit/76543d1cb2813e6d7c64a6012fcc50bd2caab92c) docs: replace nsenter with chroot in README.md
* [`f071725`](https://github.com/siderolabs/extensions/commit/f0717255a9f0c55eff1e214d8b370d99cf8103c0) docs: add steps to create mdadm software raid
* [`544ebcf`](https://github.com/siderolabs/extensions/commit/544ebcfef798f97708cb09d33f6afe2dd8986171) fix: iscsi-tools extension
* [`8e296ab`](https://github.com/siderolabs/extensions/commit/8e296ab29000e6550e7c516cd98c07e6e7f3750e) docs: fix Aenix company name
* [`c7eb377`](https://github.com/siderolabs/extensions/commit/c7eb3771943eba55ec5fe61f7a31e6fce61d73d5) feat: glibc extension
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.8.0](https://github.com/siderolabs/extensions/releases/tag/v1.8.0)

## [Talos System Extensions 1.9.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.9.0-alpha.0) (2024-10-18)

Welcome to the v1.9.0-alpha.0 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.9/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### LLDP

lldpd is now available as a system extension.


### Component Updates

Linux Firmware: 20241017
Tailscale: 1.76.0
crun: 1.17
gvisor: 20241007.0
spin: 0.16.0
ecr-credential-provider: 1.31.1


### Contributors

* Andrey Smirnov
* Dmitry Sharshakov
* Niklas Wik
* Noel Georgi

### Changes
<details><summary>8 commits</summary>
<p>

* [`862d0ac`](https://github.com/siderolabs/extensions/commit/862d0ac63d3a491c87283c9a7423f4d516e42b6f) feat: update dependencies
* [`8a7635b`](https://github.com/siderolabs/extensions/commit/8a7635b2c9ce2cb02fd389192da3d1392ca37674) feat: introduce LLDPD extension service
* [`6a184b8`](https://github.com/siderolabs/extensions/commit/6a184b8acc85237fa0e929277ad028ff576fe4b5) chore: update packages, prepare for v1.9
* [`d474848`](https://github.com/siderolabs/extensions/commit/d4748485fd6d5e44650259fa42bb228fc5cfef12) fix: zfs extensions service yaml to proper path
* [`bfaee18`](https://github.com/siderolabs/extensions/commit/bfaee18773163c35d8a8c91232fb22eaa0732619) feat: update pkgs/Linux firmware
* [`11f48c5`](https://github.com/siderolabs/extensions/commit/11f48c567ff2d491f8dd27897befcdaf87d7c989) fix: image reproducibility with finalize
* [`39d2f20`](https://github.com/siderolabs/extensions/commit/39d2f209d142bf88745aadae74df0830cd865ea4) feat: update Intel u-code to 20240910
* [`8a80e47`](https://github.com/siderolabs/extensions/commit/8a80e47219eb20c4cff2a71995984bdd7fed59fc) fix: reproducibility tests
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.8.0](https://github.com/siderolabs/extensions/releases/tag/v1.8.0)

## [Talos System Extensions 1.8.0-alpha.2](https://github.com/siderolabs/extensions/releases/tag/v1.8.0-alpha.2) (2024-08-30)

Welcome to the v1.8.0-alpha.2 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.8/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### CRUN Container Runtime

CRUN container runtime is now shipped as a Talos System Extension


### Gvisor Container Runtime

Gvisor now ships an additional runtime using `kvm` as the sandboxing mechanism.


### Intel Management Engine

Intel Management Engine (IME) modules is now shipped as a Talos System Extension.


### NVIDIA Driver and Container Toolkit

The NVIDIA drivers and the container toolkits now ships an LTS and Production version as per https://docs.nvidia.com/datacenter/tesla/drivers/index.html#lifecycle.

The new extensions are named below:

* nvidia-container-toolkit-production
* nvidia-container-toolkit-lts
* nvidia-open-gpu-kernel-modules-production
* nvidia-open-gpu-kernel-modules-lts
* nonfree-kmod-nvidia-lts
* nonfree-kmod-nvidia-production

The extensions would ship the latest version of LTS/Production drivers available at the time of Talos release.

Image Factory using an existing schematic id would upgrade the NVIDIA driver and container toolkit to the LTS version.

If production version is required, the schematic id should be updated to the production version.


### Component Updates

ZFS: 2.2.4
DRBD: 9.2.11
gasket: 5815ee3
Tailscale: 1.70.0
ecr-credential-provider: 1.31.0
qemu-guest-agent: 9.0.2
mdadm: 4.3
Intel microcode: 20240813
Linux firmware: 20240811
Spin: 0.15.1
Gvisor: 20240729.0
Wasmedge: v0.4.0
Kata Containers: 3.3.0
NVIDIA container toolkit: v1.16.1
iscsi-tools: v0.1.5
vmtoolsd: v0.6.0
util-linux-tools: 2.40.2


### Contributors

* Noel Georgi
* Andrey Smirnov
* Rui Lopes
* Bernard Gütermann
* David Peralta
* Dmitriy Matrenichev
* Henrik Gerdes
* Judah Rand
* Kingdon Barrett
* Mark S
* Markus Reiter
* Mathieu Dallaire
* Mike Beaumont
* Nick Meyer
* Sheogorath
* Sven Pfennig
* Tobias Bradtke

### Changes
<details><summary>49 commits</summary>
<p>

* [`d33d428`](https://github.com/siderolabs/extensions/commit/d33d428dbf0df0cc6845174c7a607c61556d05ee) feat: add `uinput` driver extension
* [`4563de5`](https://github.com/siderolabs/extensions/commit/4563de58b26e1a9d6990e590e2afaf26602306d2) feat: bump dependencies
* [`e753a74`](https://github.com/siderolabs/extensions/commit/e753a74ee12a1715e9d6da7ef253cb255fe40038) chore: add MAINTAINERS file
* [`bb94c9d`](https://github.com/siderolabs/extensions/commit/bb94c9d65a46caf2f2d03e191dab677d45f976f4) fix(stargz): set default root path
* [`a5a6365`](https://github.com/siderolabs/extensions/commit/a5a636538dfccbef0dfd536511a81cf24a26199a) chore: add missing license
* [`5f4947e`](https://github.com/siderolabs/extensions/commit/5f4947e28adedceb45ee610a456ad8b83cbe3240) docs: fix link to kspp page
* [`03337d7`](https://github.com/siderolabs/extensions/commit/03337d706dd6148bd7e0a0350c43031b78ff58c1) chore: bump deps
* [`cac285c`](https://github.com/siderolabs/extensions/commit/cac285cbd24e086af09df99a0b602538b70afd3a) chore: bump deps
* [`d6c324d`](https://github.com/siderolabs/extensions/commit/d6c324dc1f65cd2f0ffa25f6b7e24991923fc1c2) chore: bump deps
* [`37f2297`](https://github.com/siderolabs/extensions/commit/37f2297e6bdcc8fdc1eb5efcf166dfd4534bf261) feat: support lts and production nvidia modules
* [`6e6f029`](https://github.com/siderolabs/extensions/commit/6e6f0293e138ebc1d3e4a15106614a46d386488c) docs: update README and release notes
* [`26c505d`](https://github.com/siderolabs/extensions/commit/26c505db8cd83dd5aa7534440b7f467a6b4fa58e) feat: add crun container-runtime extension
* [`c002fba`](https://github.com/siderolabs/extensions/commit/c002fbaf4853f433b4b86598311e74e2c87a4974) feat(mei): add extension to provide Intel Management Engine drivers
* [`ab77645`](https://github.com/siderolabs/extensions/commit/ab77645a00cb074e2b52338540bca9a0cca72a6f) fix: update CRI config parts for containerd config v3
* [`c536209`](https://github.com/siderolabs/extensions/commit/c536209ef820ab59b5a653fa451027f82c433352) feat(gvisor): add new runtime class with kvm support
* [`b48d3a6`](https://github.com/siderolabs/extensions/commit/b48d3a65e6e9f57d0a54a58a1edbc8a735d900c8) chore: update extensions validator
* [`807f599`](https://github.com/siderolabs/extensions/commit/807f59946ce986089fb2664eca31639ec6531302) release(v1.8.0-alpha.1): prepare release
* [`d6773dd`](https://github.com/siderolabs/extensions/commit/d6773dd25aba4f101c2f49b251ea0a67ba242869) chore: bump deps
* [`86511df`](https://github.com/siderolabs/extensions/commit/86511dff5bea2964987bd750c31ddba3bf3e214a) chore: update spin extension to v0.15.0
* [`5334e89`](https://github.com/siderolabs/extensions/commit/5334e89374e6fb0766e25b0e702647908f4abfc0) fix: glibc search paths for nvidia
* [`3197e22`](https://github.com/siderolabs/extensions/commit/3197e22a3121f71dde52a78792f67962696496b9) docs: improve `nut-client` docs
* [`f0b6082`](https://github.com/siderolabs/extensions/commit/f0b6082466dc78a309d1e9a7d8525497d714d4d4) chore: bump tailscale to v1.68.1
* [`8e946d6`](https://github.com/siderolabs/extensions/commit/8e946d688f52e689fa447bc0daecf09fe84623b0) fix: tgtd mount propagation
* [`5904e12`](https://github.com/siderolabs/extensions/commit/5904e12cec3312c4808e9c65c94a6c555d17caa3) chore: add cache paths for go builds
* [`b840088`](https://github.com/siderolabs/extensions/commit/b84008881c27a419e8153aa4d8332a5a59717734) chore: drop `/proc` mounts
* [`3526f45`](https://github.com/siderolabs/extensions/commit/3526f4507a568bc2d6101ab1d15c9b29ddea47eb) fix: zfs extensions with nvidia
* [`13f56fc`](https://github.com/siderolabs/extensions/commit/13f56fcac088dc8ba61198e15c633781d2e6ee20) chore: rename talos-vmtoolsd -> vmtoolsd-guest-agent
* [`5e92a6c`](https://github.com/siderolabs/extensions/commit/5e92a6cb93f14fbef5230816b98d45f6366ab020) chore: use fedora mirror for glibc
* [`4ed9ee5`](https://github.com/siderolabs/extensions/commit/4ed9ee584987bf47e20246680081d589b664a413) fix: zfs-tools libtirpc path
* [`cce3b41`](https://github.com/siderolabs/extensions/commit/cce3b415e07a471d53d41b188a7c325ccc6a4d27) fix: gvisor-debug extension
* [`d07caf7`](https://github.com/siderolabs/extensions/commit/d07caf7eed782732f09427f863d23e0dffe9c034) chore: add extensions validator
* [`d1a0ce8`](https://github.com/siderolabs/extensions/commit/d1a0ce88c4e25e63cbe6e9664c621b75dea505bd) feat: include nsenter when building util-linux binaries
* [`8abfa20`](https://github.com/siderolabs/extensions/commit/8abfa2085a1737b32a61ee7c6d20e93de3dd3d94) chore(ci): drop drone
* [`7f39bce`](https://github.com/siderolabs/extensions/commit/7f39bceabb076b9157cb19c335956f51d5ad6849) feat: update Linux firmware to 20240513
* [`44a6ab1`](https://github.com/siderolabs/extensions/commit/44a6ab1ec6fdf97c1902e92c330fc75bb8e52a93) feat: update Intel u-code to 20240514
* [`d6f0b54`](https://github.com/siderolabs/extensions/commit/d6f0b546612bb9a3cf5fc0d1689d59a8308b1259) chore: update spin extension to v0.14.1
* [`01808ff`](https://github.com/siderolabs/extensions/commit/01808ff2feef6d4bd29bfde10dca102219d5b2f7) feat: update mdadm to 4.3
* [`2f97116`](https://github.com/siderolabs/extensions/commit/2f97116a50ee6b8cb6f1dd44f53a3db031a35711) feat: update dependencies
* [`dffe8b9`](https://github.com/siderolabs/extensions/commit/dffe8b9546a4ca20640c3754460b457814db16f2) fix: extension name in manifests
* [`d21bc48`](https://github.com/siderolabs/extensions/commit/d21bc482678c030957cff4ada882afcd372f8ab5) feat: update zfs extension to v2.2.4
* [`80c5113`](https://github.com/siderolabs/extensions/commit/80c5113abfabdd868a069410bcd86cdeec790b48) release(v1.8.0-alpha.0): prepare release
* [`dd85754`](https://github.com/siderolabs/extensions/commit/dd8575455e1aaf2eff31f0217bffc15d9d6450d3) fix: rekres the repo after pkgs bump
* [`d589ad0`](https://github.com/siderolabs/extensions/commit/d589ad0d7955a3437063dcaabc21d9b68816eff1) release(v1.8.0-alpha.0): prepare release
* [`882b4ac`](https://github.com/siderolabs/extensions/commit/882b4ac9f59d0daf4e5bbd5d15ae80861a778de3) fix: version in util-linux manifest
* [`4df073a`](https://github.com/siderolabs/extensions/commit/4df073ab7f5cc5d037a466e048144b57d655e408) chore(ci): only build amd64 for extensions e2e
* [`69fe96c`](https://github.com/siderolabs/extensions/commit/69fe96ccc330ac84bcbc265a4a95bd35cfe0df2a) docs: improve ExtensionServiceConfig docs
* [`8672c3b`](https://github.com/siderolabs/extensions/commit/8672c3baf51f36a9e796b0532db0f1159847cf4d) chore: update pkgs version to the latest
* [`76d3797`](https://github.com/siderolabs/extensions/commit/76d3797fedad56cbe1c0a9ba85328ac9f545ce35) docs: update Spin README.md
* [`213ef32`](https://github.com/siderolabs/extensions/commit/213ef326c12bcf0a97f8fd29caa46ff40c96a310) feat: add spin wasm runtime
</p>
</details>

### Changes since v1.8.0-alpha.1
<details><summary>15 commits</summary>
<p>

* [`d33d428`](https://github.com/siderolabs/extensions/commit/d33d428dbf0df0cc6845174c7a607c61556d05ee) feat: add `uinput` driver extension
* [`4563de5`](https://github.com/siderolabs/extensions/commit/4563de58b26e1a9d6990e590e2afaf26602306d2) feat: bump dependencies
* [`e753a74`](https://github.com/siderolabs/extensions/commit/e753a74ee12a1715e9d6da7ef253cb255fe40038) chore: add MAINTAINERS file
* [`bb94c9d`](https://github.com/siderolabs/extensions/commit/bb94c9d65a46caf2f2d03e191dab677d45f976f4) fix(stargz): set default root path
* [`a5a6365`](https://github.com/siderolabs/extensions/commit/a5a636538dfccbef0dfd536511a81cf24a26199a) chore: add missing license
* [`5f4947e`](https://github.com/siderolabs/extensions/commit/5f4947e28adedceb45ee610a456ad8b83cbe3240) docs: fix link to kspp page
* [`03337d7`](https://github.com/siderolabs/extensions/commit/03337d706dd6148bd7e0a0350c43031b78ff58c1) chore: bump deps
* [`cac285c`](https://github.com/siderolabs/extensions/commit/cac285cbd24e086af09df99a0b602538b70afd3a) chore: bump deps
* [`d6c324d`](https://github.com/siderolabs/extensions/commit/d6c324dc1f65cd2f0ffa25f6b7e24991923fc1c2) chore: bump deps
* [`37f2297`](https://github.com/siderolabs/extensions/commit/37f2297e6bdcc8fdc1eb5efcf166dfd4534bf261) feat: support lts and production nvidia modules
* [`6e6f029`](https://github.com/siderolabs/extensions/commit/6e6f0293e138ebc1d3e4a15106614a46d386488c) docs: update README and release notes
* [`26c505d`](https://github.com/siderolabs/extensions/commit/26c505db8cd83dd5aa7534440b7f467a6b4fa58e) feat: add crun container-runtime extension
* [`c002fba`](https://github.com/siderolabs/extensions/commit/c002fbaf4853f433b4b86598311e74e2c87a4974) feat(mei): add extension to provide Intel Management Engine drivers
* [`ab77645`](https://github.com/siderolabs/extensions/commit/ab77645a00cb074e2b52338540bca9a0cca72a6f) fix: update CRI config parts for containerd config v3
* [`c536209`](https://github.com/siderolabs/extensions/commit/c536209ef820ab59b5a653fa451027f82c433352) feat(gvisor): add new runtime class with kvm support
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.7.0](https://github.com/siderolabs/extensions/releases/tag/v1.7.0)

## [Talos System Extensions 1.8.0-alpha.1](https://github.com/siderolabs/extensions/releases/tag/v1.8.0-alpha.1) (2024-07-08)

Welcome to the v1.8.0-alpha.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.8/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

ZFS: 2.2.4
DRBD: 9.2.9
gasket: 5815ee3
Tailscale: 1.64.2
ecr-credential-provider: 1.30.2
qemu-guest-agent: 9.0.1
mdadm: 4.3
Intel microcode: 20240531
Linux firmware: 20240513
Spin: 1.5.0
Gvisor: 20240624.0
Wasmedge: v0.4.0
Kata Containers: 3.6.0
NVIDIA container toolkit: v1.15.0
iscsi-tools: v0.1.5


### Contributors

* Noel Georgi
* Andrey Smirnov
* Rui Lopes
* Bernard Gütermann
* Kingdon Barrett
* Mark S
* Markus Reiter
* Mathieu Dallaire
* Mike Beaumont
* Sven Pfennig

### Changes
<details><summary>32 commits</summary>
<p>

* [`d6773dd`](https://github.com/siderolabs/extensions/commit/d6773dd25aba4f101c2f49b251ea0a67ba242869) chore: bump deps
* [`86511df`](https://github.com/siderolabs/extensions/commit/86511dff5bea2964987bd750c31ddba3bf3e214a) chore: update spin extension to v0.15.0
* [`5334e89`](https://github.com/siderolabs/extensions/commit/5334e89374e6fb0766e25b0e702647908f4abfc0) fix: glibc search paths for nvidia
* [`3197e22`](https://github.com/siderolabs/extensions/commit/3197e22a3121f71dde52a78792f67962696496b9) docs: improve `nut-client` docs
* [`f0b6082`](https://github.com/siderolabs/extensions/commit/f0b6082466dc78a309d1e9a7d8525497d714d4d4) chore: bump tailscale to v1.68.1
* [`8e946d6`](https://github.com/siderolabs/extensions/commit/8e946d688f52e689fa447bc0daecf09fe84623b0) fix: tgtd mount propagation
* [`5904e12`](https://github.com/siderolabs/extensions/commit/5904e12cec3312c4808e9c65c94a6c555d17caa3) chore: add cache paths for go builds
* [`b840088`](https://github.com/siderolabs/extensions/commit/b84008881c27a419e8153aa4d8332a5a59717734) chore: drop `/proc` mounts
* [`3526f45`](https://github.com/siderolabs/extensions/commit/3526f4507a568bc2d6101ab1d15c9b29ddea47eb) fix: zfs extensions with nvidia
* [`13f56fc`](https://github.com/siderolabs/extensions/commit/13f56fcac088dc8ba61198e15c633781d2e6ee20) chore: rename talos-vmtoolsd -> vmtoolsd-guest-agent
* [`5e92a6c`](https://github.com/siderolabs/extensions/commit/5e92a6cb93f14fbef5230816b98d45f6366ab020) chore: use fedora mirror for glibc
* [`4ed9ee5`](https://github.com/siderolabs/extensions/commit/4ed9ee584987bf47e20246680081d589b664a413) fix: zfs-tools libtirpc path
* [`cce3b41`](https://github.com/siderolabs/extensions/commit/cce3b415e07a471d53d41b188a7c325ccc6a4d27) fix: gvisor-debug extension
* [`d07caf7`](https://github.com/siderolabs/extensions/commit/d07caf7eed782732f09427f863d23e0dffe9c034) chore: add extensions validator
* [`d1a0ce8`](https://github.com/siderolabs/extensions/commit/d1a0ce88c4e25e63cbe6e9664c621b75dea505bd) feat: include nsenter when building util-linux binaries
* [`8abfa20`](https://github.com/siderolabs/extensions/commit/8abfa2085a1737b32a61ee7c6d20e93de3dd3d94) chore(ci): drop drone
* [`7f39bce`](https://github.com/siderolabs/extensions/commit/7f39bceabb076b9157cb19c335956f51d5ad6849) feat: update Linux firmware to 20240513
* [`44a6ab1`](https://github.com/siderolabs/extensions/commit/44a6ab1ec6fdf97c1902e92c330fc75bb8e52a93) feat: update Intel u-code to 20240514
* [`d6f0b54`](https://github.com/siderolabs/extensions/commit/d6f0b546612bb9a3cf5fc0d1689d59a8308b1259) chore: update spin extension to v0.14.1
* [`01808ff`](https://github.com/siderolabs/extensions/commit/01808ff2feef6d4bd29bfde10dca102219d5b2f7) feat: update mdadm to 4.3
* [`2f97116`](https://github.com/siderolabs/extensions/commit/2f97116a50ee6b8cb6f1dd44f53a3db031a35711) feat: update dependencies
* [`dffe8b9`](https://github.com/siderolabs/extensions/commit/dffe8b9546a4ca20640c3754460b457814db16f2) fix: extension name in manifests
* [`d21bc48`](https://github.com/siderolabs/extensions/commit/d21bc482678c030957cff4ada882afcd372f8ab5) feat: update zfs extension to v2.2.4
* [`80c5113`](https://github.com/siderolabs/extensions/commit/80c5113abfabdd868a069410bcd86cdeec790b48) release(v1.8.0-alpha.0): prepare release
* [`dd85754`](https://github.com/siderolabs/extensions/commit/dd8575455e1aaf2eff31f0217bffc15d9d6450d3) fix: rekres the repo after pkgs bump
* [`d589ad0`](https://github.com/siderolabs/extensions/commit/d589ad0d7955a3437063dcaabc21d9b68816eff1) release(v1.8.0-alpha.0): prepare release
* [`882b4ac`](https://github.com/siderolabs/extensions/commit/882b4ac9f59d0daf4e5bbd5d15ae80861a778de3) fix: version in util-linux manifest
* [`4df073a`](https://github.com/siderolabs/extensions/commit/4df073ab7f5cc5d037a466e048144b57d655e408) chore(ci): only build amd64 for extensions e2e
* [`69fe96c`](https://github.com/siderolabs/extensions/commit/69fe96ccc330ac84bcbc265a4a95bd35cfe0df2a) docs: improve ExtensionServiceConfig docs
* [`8672c3b`](https://github.com/siderolabs/extensions/commit/8672c3baf51f36a9e796b0532db0f1159847cf4d) chore: update pkgs version to the latest
* [`76d3797`](https://github.com/siderolabs/extensions/commit/76d3797fedad56cbe1c0a9ba85328ac9f545ce35) docs: update Spin README.md
* [`213ef32`](https://github.com/siderolabs/extensions/commit/213ef326c12bcf0a97f8fd29caa46ff40c96a310) feat: add spin wasm runtime
</p>
</details>

### Changes since v1.8.0-alpha.0
<details><summary>23 commits</summary>
<p>

* [`d6773dd`](https://github.com/siderolabs/extensions/commit/d6773dd25aba4f101c2f49b251ea0a67ba242869) chore: bump deps
* [`86511df`](https://github.com/siderolabs/extensions/commit/86511dff5bea2964987bd750c31ddba3bf3e214a) chore: update spin extension to v0.15.0
* [`5334e89`](https://github.com/siderolabs/extensions/commit/5334e89374e6fb0766e25b0e702647908f4abfc0) fix: glibc search paths for nvidia
* [`3197e22`](https://github.com/siderolabs/extensions/commit/3197e22a3121f71dde52a78792f67962696496b9) docs: improve `nut-client` docs
* [`f0b6082`](https://github.com/siderolabs/extensions/commit/f0b6082466dc78a309d1e9a7d8525497d714d4d4) chore: bump tailscale to v1.68.1
* [`8e946d6`](https://github.com/siderolabs/extensions/commit/8e946d688f52e689fa447bc0daecf09fe84623b0) fix: tgtd mount propagation
* [`5904e12`](https://github.com/siderolabs/extensions/commit/5904e12cec3312c4808e9c65c94a6c555d17caa3) chore: add cache paths for go builds
* [`b840088`](https://github.com/siderolabs/extensions/commit/b84008881c27a419e8153aa4d8332a5a59717734) chore: drop `/proc` mounts
* [`3526f45`](https://github.com/siderolabs/extensions/commit/3526f4507a568bc2d6101ab1d15c9b29ddea47eb) fix: zfs extensions with nvidia
* [`13f56fc`](https://github.com/siderolabs/extensions/commit/13f56fcac088dc8ba61198e15c633781d2e6ee20) chore: rename talos-vmtoolsd -> vmtoolsd-guest-agent
* [`5e92a6c`](https://github.com/siderolabs/extensions/commit/5e92a6cb93f14fbef5230816b98d45f6366ab020) chore: use fedora mirror for glibc
* [`4ed9ee5`](https://github.com/siderolabs/extensions/commit/4ed9ee584987bf47e20246680081d589b664a413) fix: zfs-tools libtirpc path
* [`cce3b41`](https://github.com/siderolabs/extensions/commit/cce3b415e07a471d53d41b188a7c325ccc6a4d27) fix: gvisor-debug extension
* [`d07caf7`](https://github.com/siderolabs/extensions/commit/d07caf7eed782732f09427f863d23e0dffe9c034) chore: add extensions validator
* [`d1a0ce8`](https://github.com/siderolabs/extensions/commit/d1a0ce88c4e25e63cbe6e9664c621b75dea505bd) feat: include nsenter when building util-linux binaries
* [`8abfa20`](https://github.com/siderolabs/extensions/commit/8abfa2085a1737b32a61ee7c6d20e93de3dd3d94) chore(ci): drop drone
* [`7f39bce`](https://github.com/siderolabs/extensions/commit/7f39bceabb076b9157cb19c335956f51d5ad6849) feat: update Linux firmware to 20240513
* [`44a6ab1`](https://github.com/siderolabs/extensions/commit/44a6ab1ec6fdf97c1902e92c330fc75bb8e52a93) feat: update Intel u-code to 20240514
* [`d6f0b54`](https://github.com/siderolabs/extensions/commit/d6f0b546612bb9a3cf5fc0d1689d59a8308b1259) chore: update spin extension to v0.14.1
* [`01808ff`](https://github.com/siderolabs/extensions/commit/01808ff2feef6d4bd29bfde10dca102219d5b2f7) feat: update mdadm to 4.3
* [`2f97116`](https://github.com/siderolabs/extensions/commit/2f97116a50ee6b8cb6f1dd44f53a3db031a35711) feat: update dependencies
* [`dffe8b9`](https://github.com/siderolabs/extensions/commit/dffe8b9546a4ca20640c3754460b457814db16f2) fix: extension name in manifests
* [`d21bc48`](https://github.com/siderolabs/extensions/commit/d21bc482678c030957cff4ada882afcd372f8ab5) feat: update zfs extension to v2.2.4
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.7.0](https://github.com/siderolabs/extensions/releases/tag/v1.7.0)

## [Talos System Extensions 1.8.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.8.0-alpha.0) (2024-05-02)

Welcome to the v1.8.0-alpha.0 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.8/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Contributors

* Andrey Smirnov
* Kingdon Barrett
* Mathieu Dallaire
* Noel Georgi
* Sven Pfennig

### Changes
<details><summary>8 commits</summary>
<p>

* [`dd85754`](https://github.com/siderolabs/extensions/commit/dd8575455e1aaf2eff31f0217bffc15d9d6450d3) fix: rekres the repo after pkgs bump
* [`d589ad0`](https://github.com/siderolabs/extensions/commit/d589ad0d7955a3437063dcaabc21d9b68816eff1) release(v1.8.0-alpha.0): prepare release
* [`882b4ac`](https://github.com/siderolabs/extensions/commit/882b4ac9f59d0daf4e5bbd5d15ae80861a778de3) fix: version in util-linux manifest
* [`4df073a`](https://github.com/siderolabs/extensions/commit/4df073ab7f5cc5d037a466e048144b57d655e408) chore(ci): only build amd64 for extensions e2e
* [`69fe96c`](https://github.com/siderolabs/extensions/commit/69fe96ccc330ac84bcbc265a4a95bd35cfe0df2a) docs: improve ExtensionServiceConfig docs
* [`8672c3b`](https://github.com/siderolabs/extensions/commit/8672c3baf51f36a9e796b0532db0f1159847cf4d) chore: update pkgs version to the latest
* [`76d3797`](https://github.com/siderolabs/extensions/commit/76d3797fedad56cbe1c0a9ba85328ac9f545ce35) docs: update Spin README.md
* [`213ef32`](https://github.com/siderolabs/extensions/commit/213ef326c12bcf0a97f8fd29caa46ff40c96a310) feat: add spin wasm runtime
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.7.0](https://github.com/siderolabs/extensions/releases/tag/v1.7.0)

## [Talos System Extensions 1.8.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.8.0-alpha.0) (2024-05-01)

Welcome to the v1.8.0-alpha.0 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.8/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Contributors

* Andrey Smirnov
* Kingdon Barrett
* Mathieu Dallaire
* Noel Georgi
* Sven Pfennig

### Changes
<details><summary>6 commits</summary>
<p>

* [`882b4ac`](https://github.com/siderolabs/extensions/commit/882b4ac9f59d0daf4e5bbd5d15ae80861a778de3) fix: version in util-linux manifest
* [`4df073a`](https://github.com/siderolabs/extensions/commit/4df073ab7f5cc5d037a466e048144b57d655e408) chore(ci): only build amd64 for extensions e2e
* [`69fe96c`](https://github.com/siderolabs/extensions/commit/69fe96ccc330ac84bcbc265a4a95bd35cfe0df2a) docs: improve ExtensionServiceConfig docs
* [`8672c3b`](https://github.com/siderolabs/extensions/commit/8672c3baf51f36a9e796b0532db0f1159847cf4d) chore: update pkgs version to the latest
* [`76d3797`](https://github.com/siderolabs/extensions/commit/76d3797fedad56cbe1c0a9ba85328ac9f545ce35) docs: update Spin README.md
* [`213ef32`](https://github.com/siderolabs/extensions/commit/213ef326c12bcf0a97f8fd29caa46ff40c96a310) feat: add spin wasm runtime
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.7.0](https://github.com/siderolabs/extensions/releases/tag/v1.7.0)

## [Talos System Extensions 1.7.0-alpha.1](https://github.com/siderolabs/extensions/releases/tag/v1.7.0-alpha.1) (2024-03-14)

Welcome to the v1.7.0-alpha.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.7/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

* ZFS: 2.2.3
* Linux Firmware: 20240312
* DRBD: 9.2.8
* gvisor: 20240305.0
* QEMU: 8.2.2
* Tailscale: 1.60.1
* nvidia-container-runtime: v1.14.5
* libnvidia-container: v1.14.5


### Xen Guest Agent

Xen guest agent extension is now available. This removes the previous `xe-guest-utilities` extension.
See [this](https://github.com/xenserver/xe-guest-utilities/issues/118) for more info.


### Contributors

* Andrey Smirnov
* Noel Georgi
* Artem Chernyshev
* Anthony ARNAUD
* Damir Nugmanov
* Fabiano Fidêncio
* Jacob McSwain
* Miguel Angel Jaimes Linares
* Nathan Lee
* Saiyam Pathak
* Victor Seva
* j3rwin

### Changes
<details><summary>28 commits</summary>
<p>

* [`1560ee6`](https://github.com/siderolabs/extensions/commit/1560ee6f46e88c701db478ba23f718a9c8352bd1) chore: drop Drone requests for resources
* [`1459bc7`](https://github.com/siderolabs/extensions/commit/1459bc77271f5b0cc3cd90ec6c995fd1fd914be8) feat: update dependencies
* [`e5488bd`](https://github.com/siderolabs/extensions/commit/e5488bdcd36068d9a91be060ecdcbb78ca17d5bf) feat: add the new rust-based xen-guest-agent as extension
* [`ee6071c`](https://github.com/siderolabs/extensions/commit/ee6071ca7c8c9f38c3faacd872099d029ff6b87b) fix: glib build
* [`ba40f6e`](https://github.com/siderolabs/extensions/commit/ba40f6e50846def5b0620f562fc48eb111894389) feat: update Go to 1.22.1, update releases
* [`884722f`](https://github.com/siderolabs/extensions/commit/884722fdd2f2169e7ebe0e6bbad2dacb96826946) test: add gvisor-debug extension
* [`0cf50cd`](https://github.com/siderolabs/extensions/commit/0cf50cdf429d4f65c07034c4e355e49076f7eaf0) feat: update gvisor to 20240212.0
* [`5cb79a9`](https://github.com/siderolabs/extensions/commit/5cb79a9c29faa28b46770c692b54c01c78ec247d) feat: update pkgs, re-enable ZFS
* [`ab5f8d9`](https://github.com/siderolabs/extensions/commit/ab5f8d95c770c8dc6fe23af138d1b646ea4745f5) feat: add mdmon to mdadm extension
* [`9cdf805`](https://github.com/siderolabs/extensions/commit/9cdf805a5d3c5c0431a2db6f0515fe9744808232) chore: bump dependencies
* [`ec1cf8c`](https://github.com/siderolabs/extensions/commit/ec1cf8c9579b6a869a6a47e574d985233379f1d5) feat: add kata-containers extension
* [`97e59f8`](https://github.com/siderolabs/extensions/commit/97e59f833b77c0465460d7f5baf87359940cdf79) feat: realtek firmware
* [`c677b87`](https://github.com/siderolabs/extensions/commit/c677b87c58a5768fd3e294b8c4efe9a6f51ee6f3) feat: use `ExtensionServiceConfig` document
* [`fbaefd5`](https://github.com/siderolabs/extensions/commit/fbaefd573c86799d1053bb50bfb8e13a4ea0fb80) feat: add wasm runtime
* [`cde1e75`](https://github.com/siderolabs/extensions/commit/cde1e7529461237eddacd67f5b3c43795dc4cf92) release(v1.7.0-alpha.0): prepare release
* [`799a340`](https://github.com/siderolabs/extensions/commit/799a3403eab2e28baed1e44063c1851d72215eee) fix: set `DATABASE_PATH` for fabricmanager
* [`4a93d56`](https://github.com/siderolabs/extensions/commit/4a93d56ec714b518f871c735763e7b2c5486b04b) chore: adjust extension versions (tags) to drop Talos version
* [`66a1265`](https://github.com/siderolabs/extensions/commit/66a12656b2f08013d243c4e7f002faef341f73aa) feat: update extensions for Linux 6.6.x
* [`0273100`](https://github.com/siderolabs/extensions/commit/0273100f77efb8263cbffa8dfa6491b7edadd7dc) fix: allow to use custom PKGS_PREFIX
* [`805b20f`](https://github.com/siderolabs/extensions/commit/805b20fee819086bc13c409c8ce6d7f8e239eec8) fix: update SHAs for util-linux
* [`056e5a8`](https://github.com/siderolabs/extensions/commit/056e5a831abfa2d14b90255df5b51b9ed40a5d07) chore: bump dependencies
* [`5a97a46`](https://github.com/siderolabs/extensions/commit/5a97a46e24342ac7453c41578904e9c94a594da5) feat: usb video class (webcam) extension
* [`411dbc2`](https://github.com/siderolabs/extensions/commit/411dbc23f4a7251dd8be939396a718399fe83f8f) fix: qemu-guest-agent not starting in maintenance mode
* [`57503cc`](https://github.com/siderolabs/extensions/commit/57503ccd3bbba510401dbde67340b71eee0adfee) feat: generate extensions descriptions file as part of extensions image
* [`3104df1`](https://github.com/siderolabs/extensions/commit/3104df1f68fc5edc88ff30a64e7f9258fae9ef64) feat: generate a single extensions information file for all extensions
* [`32f106f`](https://github.com/siderolabs/extensions/commit/32f106f851278365320012eabacb6d6386833494) fix: zfs extension
* [`7ba3b3a`](https://github.com/siderolabs/extensions/commit/7ba3b3a8b5c755aa3936804dc6d04d5173f5a115) feat: extension providing QLogic firmware
* [`622ec82`](https://github.com/siderolabs/extensions/commit/622ec82db2121ea353d7bd88084e29883d446b52) chore: pull in latest pkgs
</p>
</details>

### Changes since v1.7.0-alpha.0
<details><summary>14 commits</summary>
<p>

* [`1560ee6`](https://github.com/siderolabs/extensions/commit/1560ee6f46e88c701db478ba23f718a9c8352bd1) chore: drop Drone requests for resources
* [`1459bc7`](https://github.com/siderolabs/extensions/commit/1459bc77271f5b0cc3cd90ec6c995fd1fd914be8) feat: update dependencies
* [`e5488bd`](https://github.com/siderolabs/extensions/commit/e5488bdcd36068d9a91be060ecdcbb78ca17d5bf) feat: add the new rust-based xen-guest-agent as extension
* [`ee6071c`](https://github.com/siderolabs/extensions/commit/ee6071ca7c8c9f38c3faacd872099d029ff6b87b) fix: glib build
* [`ba40f6e`](https://github.com/siderolabs/extensions/commit/ba40f6e50846def5b0620f562fc48eb111894389) feat: update Go to 1.22.1, update releases
* [`884722f`](https://github.com/siderolabs/extensions/commit/884722fdd2f2169e7ebe0e6bbad2dacb96826946) test: add gvisor-debug extension
* [`0cf50cd`](https://github.com/siderolabs/extensions/commit/0cf50cdf429d4f65c07034c4e355e49076f7eaf0) feat: update gvisor to 20240212.0
* [`5cb79a9`](https://github.com/siderolabs/extensions/commit/5cb79a9c29faa28b46770c692b54c01c78ec247d) feat: update pkgs, re-enable ZFS
* [`ab5f8d9`](https://github.com/siderolabs/extensions/commit/ab5f8d95c770c8dc6fe23af138d1b646ea4745f5) feat: add mdmon to mdadm extension
* [`9cdf805`](https://github.com/siderolabs/extensions/commit/9cdf805a5d3c5c0431a2db6f0515fe9744808232) chore: bump dependencies
* [`ec1cf8c`](https://github.com/siderolabs/extensions/commit/ec1cf8c9579b6a869a6a47e574d985233379f1d5) feat: add kata-containers extension
* [`97e59f8`](https://github.com/siderolabs/extensions/commit/97e59f833b77c0465460d7f5baf87359940cdf79) feat: realtek firmware
* [`c677b87`](https://github.com/siderolabs/extensions/commit/c677b87c58a5768fd3e294b8c4efe9a6f51ee6f3) feat: use `ExtensionServiceConfig` document
* [`fbaefd5`](https://github.com/siderolabs/extensions/commit/fbaefd573c86799d1053bb50bfb8e13a4ea0fb80) feat: add wasm runtime
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.6.0](https://github.com/siderolabs/extensions/releases/tag/v1.6.0)

## [Talos System Extensions 1.7.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.7.0-alpha.0) (2024-02-01)

Welcome to the v1.7.0-alpha.0 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.7/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

* ZFS: 2.2.2
* Linux Firmware: 20240115
* DRBD: 9.2.7
* gvisor: 20240109.0
* QEMU: 8.2.0
* Tailscale: 1.56.1


### Contributors

* Andrey Smirnov
* Artem Chernyshev
* Noel Georgi
* Anthony ARNAUD
* Jacob McSwain
* Miguel Angel Jaimes Linares
* Nathan Lee

### Changes
<details><summary>13 commits</summary>
<p>

* [`799a340`](https://github.com/siderolabs/extensions/commit/799a3403eab2e28baed1e44063c1851d72215eee) fix: set `DATABASE_PATH` for fabricmanager
* [`4a93d56`](https://github.com/siderolabs/extensions/commit/4a93d56ec714b518f871c735763e7b2c5486b04b) chore: adjust extension versions (tags) to drop Talos version
* [`66a1265`](https://github.com/siderolabs/extensions/commit/66a12656b2f08013d243c4e7f002faef341f73aa) feat: update extensions for Linux 6.6.x
* [`0273100`](https://github.com/siderolabs/extensions/commit/0273100f77efb8263cbffa8dfa6491b7edadd7dc) fix: allow to use custom PKGS_PREFIX
* [`805b20f`](https://github.com/siderolabs/extensions/commit/805b20fee819086bc13c409c8ce6d7f8e239eec8) fix: update SHAs for util-linux
* [`056e5a8`](https://github.com/siderolabs/extensions/commit/056e5a831abfa2d14b90255df5b51b9ed40a5d07) chore: bump dependencies
* [`5a97a46`](https://github.com/siderolabs/extensions/commit/5a97a46e24342ac7453c41578904e9c94a594da5) feat: usb video class (webcam) extension
* [`411dbc2`](https://github.com/siderolabs/extensions/commit/411dbc23f4a7251dd8be939396a718399fe83f8f) fix: qemu-guest-agent not starting in maintenance mode
* [`57503cc`](https://github.com/siderolabs/extensions/commit/57503ccd3bbba510401dbde67340b71eee0adfee) feat: generate extensions descriptions file as part of extensions image
* [`3104df1`](https://github.com/siderolabs/extensions/commit/3104df1f68fc5edc88ff30a64e7f9258fae9ef64) feat: generate a single extensions information file for all extensions
* [`32f106f`](https://github.com/siderolabs/extensions/commit/32f106f851278365320012eabacb6d6386833494) fix: zfs extension
* [`7ba3b3a`](https://github.com/siderolabs/extensions/commit/7ba3b3a8b5c755aa3936804dc6d04d5173f5a115) feat: extension providing QLogic firmware
* [`622ec82`](https://github.com/siderolabs/extensions/commit/622ec82db2121ea353d7bd88084e29883d446b52) chore: pull in latest pkgs
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.6.0](https://github.com/siderolabs/extensions/releases/tag/v1.6.0)

## [Talos System Extensions 1.6.0-alpha.2](https://github.com/siderolabs/extensions/releases/tag/v1.6.0-alpha.2) (2023-11-21)

Welcome to the v1.6.0-alpha.2 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.6/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### AMD GPU firmware

AMD GPU firmware is now shipped as a system extension.


### binfmt_misc

`binfmt_misc` Linux module extension is now supported as Talos System Extension.


### BTRFS

BTRFS drivers is now supported as Talos System Extension.


### Intel ICE firmware

Intel ice firmware (E810) is now shipped as a system extension.


### stargz-snapshotter

stargz-snapshotter extension is now supported as Talos System Extension.


### Component Updates

* DRBD: 9.2.6
* QEMU agent: v8.1.1
* Tailscale: 1.50.0
* Xen Guest Utilities: 8.3.1
* Linux Firmware: 20231030
* gasket driver: 09385d4
* Intel ucode: 20231114


### Util Linux Tools

Util Linux Tools is now shipped as a system extension. Includes fstrim only.


### ZFS

ZFS extensions now also ship zfs-tools and an extension service that imports all zpools on startup.


### Contributors

* Noel Georgi
* Andrey Smirnov
* Andrew Rynhard
* Spencer Smith
* Cas de Reuver
* David Donchez
* Enno Boland
* Serge Logvinov
* Ströger Florian
* Victor Seva
* Will Glynn
* cDR (Taco)

### Changes
<details><summary>36 commits</summary>
<p>

* [`9a57c65`](https://github.com/siderolabs/extensions/commit/9a57c65d10068b9d28230a34d60b22ecfdbf1664) feat: bump dependencies
* [`af706ee`](https://github.com/siderolabs/extensions/commit/af706eedd4e6b1c38cca86fbcbc7c4b1c397d3a9) chore: bump Intel ucode to 20231114 (CVE-2023-23583)
* [`f990683`](https://github.com/siderolabs/extensions/commit/f990683ef02db00219af69071bd4486b20a8b05d) chore: drop rexec in pid 1 namespace for zpool-import
* [`2859c23`](https://github.com/siderolabs/extensions/commit/2859c2360934e019d9e0e93a0bde98a34780f5b7) fix: iscsi-tools install `prefix`
* [`2f8e401`](https://github.com/siderolabs/extensions/commit/2f8e401d21c4d9b856429653b08da74a5d5d496e) feat: create 'ecr-credential-provider' extension
* [`5914368`](https://github.com/siderolabs/extensions/commit/591436817af26272a7168557f811e39e1bb41f8c) fix: enable rootfs propagation for stargz-snapshotter
* [`8eb47b1`](https://github.com/siderolabs/extensions/commit/8eb47b129ad4fbf1a0bdcd77e1a6c1a1653824fe) chore: update packages
* [`d2aaecf`](https://github.com/siderolabs/extensions/commit/d2aaecfa1b9a4bf7d29dae478056a0f63477f281) feat: update Go to 1.21.4
* [`7c68b1b`](https://github.com/siderolabs/extensions/commit/7c68b1b932c9f9d10afe2e1efc5401ef528fbf81) chore: use kres to manage project
* [`d7fdcc9`](https://github.com/siderolabs/extensions/commit/d7fdcc958456c450c9d4fe53e2d1738b2305c25d) chore: move to gh actions
* [`280ff1f`](https://github.com/siderolabs/extensions/commit/280ff1fae3cab1f8b31caf35e15bc3ff7b1d95fa) release(v1.6.0-alpha.1): prepare release
* [`9b5c56f`](https://github.com/siderolabs/extensions/commit/9b5c56f0957bbee212f29a62fd310f47ec31fba4) chore: update stargz extension
* [`dc619c4`](https://github.com/siderolabs/extensions/commit/dc619c40d0a9d884a8c493237eaf9184a5e40b43) feat: add fuse3 extension
* [`0ba9f81`](https://github.com/siderolabs/extensions/commit/0ba9f8104350fbeb2832ecf796146f1ef242fb4d) docs: update documentation on installing extensions
* [`e812c6f`](https://github.com/siderolabs/extensions/commit/e812c6f7e1b99d67027e91a7d95b2a4637b61619) docs: update README
* [`ad30c33`](https://github.com/siderolabs/extensions/commit/ad30c330cd7885436b53cca86f613a0e0e6bb24d) feat: update releases
* [`ac8ff02`](https://github.com/siderolabs/extensions/commit/ac8ff02bf406380aa609b40760695cbde2ab9d38) chore: update README.md
* [`46a3004`](https://github.com/siderolabs/extensions/commit/46a30042937aa5dd269415dbfb6eaf13d521c242) feat: add binfmt_misc
* [`336397b`](https://github.com/siderolabs/extensions/commit/336397bccc94bd28e2fd7c2b969d36dc68f5023d) feat: add chelsio drivers and firmware
* [`1c94a09`](https://github.com/siderolabs/extensions/commit/1c94a09b1c663c912f3452fcf3c948617a0bb373) feat: add stargz-snapshotter
* [`3765417`](https://github.com/siderolabs/extensions/commit/37654176b0edec665e9c0933545118f3b93d6f59) chore: use kernel image for upstream modules
* [`8fa50a2`](https://github.com/siderolabs/extensions/commit/8fa50a2fa25f34f38f7bf4f393be0c618f5b8bb7) chore: add extensions catalog
* [`66cbf19`](https://github.com/siderolabs/extensions/commit/66cbf19c18f4e201d61dd83b48439d0a76746cfb) chore: update pkgs
* [`dc94329`](https://github.com/siderolabs/extensions/commit/dc94329fd3d235e96c5bb962f53bedea197bead7) feat: add xe-guest-utilities extension
* [`782ccc5`](https://github.com/siderolabs/extensions/commit/782ccc52f5d95acc3b11c3890629f7177bc57bdc) chore: fix renovate config
* [`a5c0b00`](https://github.com/siderolabs/extensions/commit/a5c0b0086be035b0eeb8a4e541cc9a4658dc55e5) chore: revert nvidia bumps from #220
* [`d9145f9`](https://github.com/siderolabs/extensions/commit/d9145f9b6b7816e57ebaa4d351699be273f8ac17) chore: bump deps
* [`0cba1b6`](https://github.com/siderolabs/extensions/commit/0cba1b6c1ca50ce7ab17a286251f3b688c166888) fix: util-linux pkg name
* [`89f857d`](https://github.com/siderolabs/extensions/commit/89f857dd49ce6d35820de9c2a90988981297efc8) feat: add minimal util-linux
* [`0f4a77e`](https://github.com/siderolabs/extensions/commit/0f4a77efed51c166539b9b1e699c5ce28ef8c98f) feat: add intel ice firmware
* [`d56bab2`](https://github.com/siderolabs/extensions/commit/d56bab242144f8e08d200b5788efb09110b51847) release(v1.6.0-alpha.0): prepare release
* [`ce60bd9`](https://github.com/siderolabs/extensions/commit/ce60bd90bfa7dd5a4551f76bed644e251a6b3b94) fix: btrfs extension constraint
* [`31cdc8c`](https://github.com/siderolabs/extensions/commit/31cdc8c50594aa686f95ce453876e780a1207ab2) feat: add amdgpu firmware
* [`d369468`](https://github.com/siderolabs/extensions/commit/d369468456794d84764b79092cd1092c37481e56) feat: zfs user tools and import service
* [`0284425`](https://github.com/siderolabs/extensions/commit/02844256f9b67715dead196be4a2a47f2320befc) chore: minor fixes
* [`1533478`](https://github.com/siderolabs/extensions/commit/1533478f5945b8483121080ec847386ca4488ff5) feat: add btrfs storage extension
</p>
</details>

### Changes since v1.6.0-alpha.1
<details><summary>10 commits</summary>
<p>

* [`9a57c65`](https://github.com/siderolabs/extensions/commit/9a57c65d10068b9d28230a34d60b22ecfdbf1664) feat: bump dependencies
* [`af706ee`](https://github.com/siderolabs/extensions/commit/af706eedd4e6b1c38cca86fbcbc7c4b1c397d3a9) chore: bump Intel ucode to 20231114 (CVE-2023-23583)
* [`f990683`](https://github.com/siderolabs/extensions/commit/f990683ef02db00219af69071bd4486b20a8b05d) chore: drop rexec in pid 1 namespace for zpool-import
* [`2859c23`](https://github.com/siderolabs/extensions/commit/2859c2360934e019d9e0e93a0bde98a34780f5b7) fix: iscsi-tools install `prefix`
* [`2f8e401`](https://github.com/siderolabs/extensions/commit/2f8e401d21c4d9b856429653b08da74a5d5d496e) feat: create 'ecr-credential-provider' extension
* [`5914368`](https://github.com/siderolabs/extensions/commit/591436817af26272a7168557f811e39e1bb41f8c) fix: enable rootfs propagation for stargz-snapshotter
* [`8eb47b1`](https://github.com/siderolabs/extensions/commit/8eb47b129ad4fbf1a0bdcd77e1a6c1a1653824fe) chore: update packages
* [`d2aaecf`](https://github.com/siderolabs/extensions/commit/d2aaecfa1b9a4bf7d29dae478056a0f63477f281) feat: update Go to 1.21.4
* [`7c68b1b`](https://github.com/siderolabs/extensions/commit/7c68b1b932c9f9d10afe2e1efc5401ef528fbf81) chore: use kres to manage project
* [`d7fdcc9`](https://github.com/siderolabs/extensions/commit/d7fdcc958456c450c9d4fe53e2d1738b2305c25d) chore: move to gh actions
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.5.0](https://github.com/siderolabs/extensions/releases/tag/v1.5.0)

## [Talos System Extensions 1.6.0-alpha.1](https://github.com/siderolabs/extensions/releases/tag/v1.6.0-alpha.1) (2023-10-17)

Welcome to the v1.6.0-alpha.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.6/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### AMD GPU firmware

AMD GPU firmware is now shipped as a system extension.


### binfmt_misc

`binfmt_misc` Linux module extension is now supported as Talos System Extension.


### BTRFS

BTRFS drivers is now supported as Talos System Extension.


### Intel ICE firmware

Intel ice firmware (E810) is now shipped as a system extension.


### stargz-snapshotter

stargz-snapshotter extension is now supported as Talos System Extension.


### Component Updates

* DRBD: 9.2.5
* QEMU agent: v8.1.1
* Tailscale: 1.50.0
* Xen Guest Utilities: 8.3.1
* Linux Firmware: 20230919
* gasket driver: 09385d4


### Util Linux Tools

Util Linux Tools is now shipped as a system extension. Includes fstrim only.


### ZFS

ZFS extensions now also ship zfs-tools and an extension service that imports all zpools on startup.


### Contributors

* Noel Georgi
* Andrey Smirnov
* Andrew Rynhard
* Spencer Smith
* Cas de Reuver
* David Donchez
* Enno Boland
* Serge Logvinov
* Victor Seva
* cDR (Taco)

### Changes
<details><summary>25 commits</summary>
<p>

* [`9b5c56f`](https://github.com/siderolabs/extensions/commit/9b5c56f0957bbee212f29a62fd310f47ec31fba4) chore: update stargz extension
* [`dc619c4`](https://github.com/siderolabs/extensions/commit/dc619c40d0a9d884a8c493237eaf9184a5e40b43) feat: add fuse3 extension
* [`0ba9f81`](https://github.com/siderolabs/extensions/commit/0ba9f8104350fbeb2832ecf796146f1ef242fb4d) docs: update documentation on installing extensions
* [`e812c6f`](https://github.com/siderolabs/extensions/commit/e812c6f7e1b99d67027e91a7d95b2a4637b61619) docs: update README
* [`ad30c33`](https://github.com/siderolabs/extensions/commit/ad30c330cd7885436b53cca86f613a0e0e6bb24d) feat: update releases
* [`ac8ff02`](https://github.com/siderolabs/extensions/commit/ac8ff02bf406380aa609b40760695cbde2ab9d38) chore: update README.md
* [`46a3004`](https://github.com/siderolabs/extensions/commit/46a30042937aa5dd269415dbfb6eaf13d521c242) feat: add binfmt_misc
* [`336397b`](https://github.com/siderolabs/extensions/commit/336397bccc94bd28e2fd7c2b969d36dc68f5023d) feat: add chelsio drivers and firmware
* [`1c94a09`](https://github.com/siderolabs/extensions/commit/1c94a09b1c663c912f3452fcf3c948617a0bb373) feat: add stargz-snapshotter
* [`3765417`](https://github.com/siderolabs/extensions/commit/37654176b0edec665e9c0933545118f3b93d6f59) chore: use kernel image for upstream modules
* [`8fa50a2`](https://github.com/siderolabs/extensions/commit/8fa50a2fa25f34f38f7bf4f393be0c618f5b8bb7) chore: add extensions catalog
* [`66cbf19`](https://github.com/siderolabs/extensions/commit/66cbf19c18f4e201d61dd83b48439d0a76746cfb) chore: update pkgs
* [`dc94329`](https://github.com/siderolabs/extensions/commit/dc94329fd3d235e96c5bb962f53bedea197bead7) feat: add xe-guest-utilities extension
* [`782ccc5`](https://github.com/siderolabs/extensions/commit/782ccc52f5d95acc3b11c3890629f7177bc57bdc) chore: fix renovate config
* [`a5c0b00`](https://github.com/siderolabs/extensions/commit/a5c0b0086be035b0eeb8a4e541cc9a4658dc55e5) chore: revert nvidia bumps from #220
* [`d9145f9`](https://github.com/siderolabs/extensions/commit/d9145f9b6b7816e57ebaa4d351699be273f8ac17) chore: bump deps
* [`0cba1b6`](https://github.com/siderolabs/extensions/commit/0cba1b6c1ca50ce7ab17a286251f3b688c166888) fix: util-linux pkg name
* [`89f857d`](https://github.com/siderolabs/extensions/commit/89f857dd49ce6d35820de9c2a90988981297efc8) feat: add minimal util-linux
* [`0f4a77e`](https://github.com/siderolabs/extensions/commit/0f4a77efed51c166539b9b1e699c5ce28ef8c98f) feat: add intel ice firmware
* [`d56bab2`](https://github.com/siderolabs/extensions/commit/d56bab242144f8e08d200b5788efb09110b51847) release(v1.6.0-alpha.0): prepare release
* [`ce60bd9`](https://github.com/siderolabs/extensions/commit/ce60bd90bfa7dd5a4551f76bed644e251a6b3b94) fix: btrfs extension constraint
* [`31cdc8c`](https://github.com/siderolabs/extensions/commit/31cdc8c50594aa686f95ce453876e780a1207ab2) feat: add amdgpu firmware
* [`d369468`](https://github.com/siderolabs/extensions/commit/d369468456794d84764b79092cd1092c37481e56) feat: zfs user tools and import service
* [`0284425`](https://github.com/siderolabs/extensions/commit/02844256f9b67715dead196be4a2a47f2320befc) chore: minor fixes
* [`1533478`](https://github.com/siderolabs/extensions/commit/1533478f5945b8483121080ec847386ca4488ff5) feat: add btrfs storage extension
</p>
</details>

### Changes since v1.6.0-alpha.0
<details><summary>19 commits</summary>
<p>

* [`9b5c56f`](https://github.com/siderolabs/extensions/commit/9b5c56f0957bbee212f29a62fd310f47ec31fba4) chore: update stargz extension
* [`dc619c4`](https://github.com/siderolabs/extensions/commit/dc619c40d0a9d884a8c493237eaf9184a5e40b43) feat: add fuse3 extension
* [`0ba9f81`](https://github.com/siderolabs/extensions/commit/0ba9f8104350fbeb2832ecf796146f1ef242fb4d) docs: update documentation on installing extensions
* [`e812c6f`](https://github.com/siderolabs/extensions/commit/e812c6f7e1b99d67027e91a7d95b2a4637b61619) docs: update README
* [`ad30c33`](https://github.com/siderolabs/extensions/commit/ad30c330cd7885436b53cca86f613a0e0e6bb24d) feat: update releases
* [`ac8ff02`](https://github.com/siderolabs/extensions/commit/ac8ff02bf406380aa609b40760695cbde2ab9d38) chore: update README.md
* [`46a3004`](https://github.com/siderolabs/extensions/commit/46a30042937aa5dd269415dbfb6eaf13d521c242) feat: add binfmt_misc
* [`336397b`](https://github.com/siderolabs/extensions/commit/336397bccc94bd28e2fd7c2b969d36dc68f5023d) feat: add chelsio drivers and firmware
* [`1c94a09`](https://github.com/siderolabs/extensions/commit/1c94a09b1c663c912f3452fcf3c948617a0bb373) feat: add stargz-snapshotter
* [`3765417`](https://github.com/siderolabs/extensions/commit/37654176b0edec665e9c0933545118f3b93d6f59) chore: use kernel image for upstream modules
* [`8fa50a2`](https://github.com/siderolabs/extensions/commit/8fa50a2fa25f34f38f7bf4f393be0c618f5b8bb7) chore: add extensions catalog
* [`66cbf19`](https://github.com/siderolabs/extensions/commit/66cbf19c18f4e201d61dd83b48439d0a76746cfb) chore: update pkgs
* [`dc94329`](https://github.com/siderolabs/extensions/commit/dc94329fd3d235e96c5bb962f53bedea197bead7) feat: add xe-guest-utilities extension
* [`782ccc5`](https://github.com/siderolabs/extensions/commit/782ccc52f5d95acc3b11c3890629f7177bc57bdc) chore: fix renovate config
* [`a5c0b00`](https://github.com/siderolabs/extensions/commit/a5c0b0086be035b0eeb8a4e541cc9a4658dc55e5) chore: revert nvidia bumps from #220
* [`d9145f9`](https://github.com/siderolabs/extensions/commit/d9145f9b6b7816e57ebaa4d351699be273f8ac17) chore: bump deps
* [`0cba1b6`](https://github.com/siderolabs/extensions/commit/0cba1b6c1ca50ce7ab17a286251f3b688c166888) fix: util-linux pkg name
* [`89f857d`](https://github.com/siderolabs/extensions/commit/89f857dd49ce6d35820de9c2a90988981297efc8) feat: add minimal util-linux
* [`0f4a77e`](https://github.com/siderolabs/extensions/commit/0f4a77efed51c166539b9b1e699c5ce28ef8c98f) feat: add intel ice firmware
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.5.0](https://github.com/siderolabs/extensions/releases/tag/v1.5.0)

## [Talos System Extensions 1.6.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.6.0-alpha.0) (2023-08-24)

Welcome to the v1.6.0-alpha.0 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.6/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### AMD GPU firmware

AMD GPU firmware is now shipped as a system extension.


### BTRFS

BTRFS drivers is now supported as Talos System Extension.


### Component Updates

* DRBD: 9.2.5


### ZFS

ZFS extensions now also ship zfs-tools and an extension service that imports all zpools on startup.


### Contributors

* Noel Georgi
* David Donchez
* Enno Boland
* Victor Seva

### Changes
<details><summary>5 commits</summary>
<p>

* [`ce60bd9`](https://github.com/siderolabs/extensions/commit/ce60bd90bfa7dd5a4551f76bed644e251a6b3b94) fix: btrfs extension constraint
* [`31cdc8c`](https://github.com/siderolabs/extensions/commit/31cdc8c50594aa686f95ce453876e780a1207ab2) feat: add amdgpu firmware
* [`d369468`](https://github.com/siderolabs/extensions/commit/d369468456794d84764b79092cd1092c37481e56) feat: zfs user tools and import service
* [`0284425`](https://github.com/siderolabs/extensions/commit/02844256f9b67715dead196be4a2a47f2320befc) chore: minor fixes
* [`1533478`](https://github.com/siderolabs/extensions/commit/1533478f5945b8483121080ec847386ca4488ff5) feat: add btrfs storage extension
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.5.0](https://github.com/siderolabs/extensions/releases/tag/v1.5.0)

## [Talos System Extensions 1.5.0-beta.1](https://github.com/siderolabs/extensions/releases/tag/v1.5.0-beta.1) (2023-08-09)

Welcome to the v1.5.0-beta.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.5/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### QEMU Guest Agent

QEMU Guest Agent is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Tailscale

Tailscale is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Thunderbolt

Thunderbolt drivers is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Component Updates

* Linux firmware: 20230804
* DRBD: 9.2.4
* Intel ucode: 20230808
* Open iSCSI: 2.1.9
* NVIDIA drivers: 535.54.03
* nvidia-container-toolkit: v1.13.5
* libnvidia-container: v1.13.5


### ZFS

ZFS is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Contributors

* Andrey Smirnov
* Noel Georgi
* Markus Reiter
* Roee Klinger
* Andrei Kvapil
* Andrey Smirnov
* Caleb Woodbine
* Igor Rzegocki
* beau trepp

### Changes
<details><summary>23 commits</summary>
<p>

* [`67ee9e3`](https://github.com/siderolabs/extensions/commit/67ee9e3b3ea9d34136505ea368e121d0733509a3) feat: microcode updates
* [`7ca8a0f`](https://github.com/siderolabs/extensions/commit/7ca8a0f04aa74bedaa034f1275f585d31e28768f) fix(docs): readability for Tailscale docs
* [`45ef4c0`](https://github.com/siderolabs/extensions/commit/45ef4c0019637382a16cc2a9943d9eb21bce1c0d) release(v1.5.0-beta.0): prepare release
* [`d4d42e5`](https://github.com/siderolabs/extensions/commit/d4d42e52d954998211d0f83d0a081da8384407fd) feat: use wolfi as base for nvidia
* [`fdfc633`](https://github.com/siderolabs/extensions/commit/fdfc63329598dfa28bf78445edc4da80c8f22f63) feat: update extension spec allowlist for opengl
* [`a8f88a9`](https://github.com/siderolabs/extensions/commit/a8f88a9ee4ee9f184363fbef0480b575e6dafbaa) chore: bump dependencies
* [`ed19c45`](https://github.com/siderolabs/extensions/commit/ed19c453dfc2b1edb68d35a8974a90158eb529e4) release(v1.5.0-alpha.3): prepare release
* [`1fdca1a`](https://github.com/siderolabs/extensions/commit/1fdca1a892054036c9661351d9740ce9e35eb2b3) feat: add thunderbolt/USB4 module
* [`14c9ec6`](https://github.com/siderolabs/extensions/commit/14c9ec6c9af143f663865d91b9188db0aa3f4e7b) feat: add 'zfs' extension
* [`cca830e`](https://github.com/siderolabs/extensions/commit/cca830ec2f0f29d2a0cf604dd7680bfa9d2a5d69) release(v1.5.0-alpha.2): prepare release
* [`44b53b8`](https://github.com/siderolabs/extensions/commit/44b53b8da53fa25646612ca32944089fc166597e) fix: ci blob upload
* [`6366b17`](https://github.com/siderolabs/extensions/commit/6366b171c75e678a4acf5dd791a4e906ed3ccf04) chore: use machined for shutdown calls
* [`048de33`](https://github.com/siderolabs/extensions/commit/048de338160553383e9e202478281ba14e7355bc) chore: bump dependencies
* [`b98fed7`](https://github.com/siderolabs/extensions/commit/b98fed77b3a9b7891c82a89fe94ce7cd6dc455fc) fix: use correct exit status in `nut-client` patch
* [`1934a5b`](https://github.com/siderolabs/extensions/commit/1934a5b12bfad6915bb8ff7f02c82ce6448c09ce) feat: add `qemu-guest-agent` extension
* [`6c502e1`](https://github.com/siderolabs/extensions/commit/6c502e1095448178b7ff04ebb97cde12fc9b6be8) feat: tailscale extension
* [`30f24da`](https://github.com/siderolabs/extensions/commit/30f24da0ce678b08705fbb9d95b85e44a5f1a353) release(v1.5.0-alpha.1): prepare release
* [`cfee4f6`](https://github.com/siderolabs/extensions/commit/cfee4f6d14c1130496bfdc1baa3c9b4d4958c9bf) chore: bump deps
* [`e814607`](https://github.com/siderolabs/extensions/commit/e8146076896e0b370ce1e1cb02d97d7dd559b0b9) docs: remove the belkin_sa from readme
* [`4cee34c`](https://github.com/siderolabs/extensions/commit/4cee34c5859e17264f899717bf843bbdccff4a4e) feat: add usb modem drivers
* [`f05c997`](https://github.com/siderolabs/extensions/commit/f05c9975c92a15c36f11dce87d05ffa0d7164e10) release(v1.5.0-alpha.0): prepare release
* [`0855dd7`](https://github.com/siderolabs/extensions/commit/0855dd71e7831c51ded46192845a85f4ba019c41) chore: bump dependencies
* [`c06874c`](https://github.com/siderolabs/extensions/commit/c06874c102fe9fc6ceb93a48be77434194f8940c) chore: add make targets to automate image signing
</p>
</details>

### Changes since v1.5.0-beta.0
<details><summary>2 commits</summary>
<p>

* [`67ee9e3`](https://github.com/siderolabs/extensions/commit/67ee9e3b3ea9d34136505ea368e121d0733509a3) feat: microcode updates
* [`7ca8a0f`](https://github.com/siderolabs/extensions/commit/7ca8a0f04aa74bedaa034f1275f585d31e28768f) fix(docs): readability for Tailscale docs
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.4.0](https://github.com/siderolabs/extensions/releases/tag/v1.4.0)

## [Talos System Extensions 1.5.0-beta.0](https://github.com/siderolabs/extensions/releases/tag/v1.5.0-beta.0) (2023-08-02)

Welcome to the v1.5.0-beta.0 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.5/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### QEMU Guest Agent

QEMU Guest Agent is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Tailscale

Tailscale is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Thunderbolt

Thunderbolt drivers is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Component Updates

* Linux firmware: 20230625
* DRBD: 9.2.4
* Intel ucode: 20230613
* Open iSCSI: 2.1.9
* NVIDIA drivers: 535.54.03
* nvidia-container-toolkit: v1.13.5
* libnvidia-container: v1.13.5


### ZFS

ZFS is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Contributors

* Andrey Smirnov
* Noel Georgi
* Markus Reiter
* Roee Klinger
* Andrei Kvapil
* Igor Rzegocki
* beau trepp

### Changes
<details><summary>20 commits</summary>
<p>

* [`d4d42e5`](https://github.com/siderolabs/extensions/commit/d4d42e52d954998211d0f83d0a081da8384407fd) feat: use wolfi as base for nvidia
* [`fdfc633`](https://github.com/siderolabs/extensions/commit/fdfc63329598dfa28bf78445edc4da80c8f22f63) feat: update extension spec allowlist for opengl
* [`a8f88a9`](https://github.com/siderolabs/extensions/commit/a8f88a9ee4ee9f184363fbef0480b575e6dafbaa) chore: bump dependencies
* [`ed19c45`](https://github.com/siderolabs/extensions/commit/ed19c453dfc2b1edb68d35a8974a90158eb529e4) release(v1.5.0-alpha.3): prepare release
* [`1fdca1a`](https://github.com/siderolabs/extensions/commit/1fdca1a892054036c9661351d9740ce9e35eb2b3) feat: add thunderbolt/USB4 module
* [`14c9ec6`](https://github.com/siderolabs/extensions/commit/14c9ec6c9af143f663865d91b9188db0aa3f4e7b) feat: add 'zfs' extension
* [`cca830e`](https://github.com/siderolabs/extensions/commit/cca830ec2f0f29d2a0cf604dd7680bfa9d2a5d69) release(v1.5.0-alpha.2): prepare release
* [`44b53b8`](https://github.com/siderolabs/extensions/commit/44b53b8da53fa25646612ca32944089fc166597e) fix: ci blob upload
* [`6366b17`](https://github.com/siderolabs/extensions/commit/6366b171c75e678a4acf5dd791a4e906ed3ccf04) chore: use machined for shutdown calls
* [`048de33`](https://github.com/siderolabs/extensions/commit/048de338160553383e9e202478281ba14e7355bc) chore: bump dependencies
* [`b98fed7`](https://github.com/siderolabs/extensions/commit/b98fed77b3a9b7891c82a89fe94ce7cd6dc455fc) fix: use correct exit status in `nut-client` patch
* [`1934a5b`](https://github.com/siderolabs/extensions/commit/1934a5b12bfad6915bb8ff7f02c82ce6448c09ce) feat: add `qemu-guest-agent` extension
* [`6c502e1`](https://github.com/siderolabs/extensions/commit/6c502e1095448178b7ff04ebb97cde12fc9b6be8) feat: tailscale extension
* [`30f24da`](https://github.com/siderolabs/extensions/commit/30f24da0ce678b08705fbb9d95b85e44a5f1a353) release(v1.5.0-alpha.1): prepare release
* [`cfee4f6`](https://github.com/siderolabs/extensions/commit/cfee4f6d14c1130496bfdc1baa3c9b4d4958c9bf) chore: bump deps
* [`e814607`](https://github.com/siderolabs/extensions/commit/e8146076896e0b370ce1e1cb02d97d7dd559b0b9) docs: remove the belkin_sa from readme
* [`4cee34c`](https://github.com/siderolabs/extensions/commit/4cee34c5859e17264f899717bf843bbdccff4a4e) feat: add usb modem drivers
* [`f05c997`](https://github.com/siderolabs/extensions/commit/f05c9975c92a15c36f11dce87d05ffa0d7164e10) release(v1.5.0-alpha.0): prepare release
* [`0855dd7`](https://github.com/siderolabs/extensions/commit/0855dd71e7831c51ded46192845a85f4ba019c41) chore: bump dependencies
* [`c06874c`](https://github.com/siderolabs/extensions/commit/c06874c102fe9fc6ceb93a48be77434194f8940c) chore: add make targets to automate image signing
</p>
</details>

### Changes since v1.5.0-alpha.3
<details><summary>3 commits</summary>
<p>

* [`d4d42e5`](https://github.com/siderolabs/extensions/commit/d4d42e52d954998211d0f83d0a081da8384407fd) feat: use wolfi as base for nvidia
* [`fdfc633`](https://github.com/siderolabs/extensions/commit/fdfc63329598dfa28bf78445edc4da80c8f22f63) feat: update extension spec allowlist for opengl
* [`a8f88a9`](https://github.com/siderolabs/extensions/commit/a8f88a9ee4ee9f184363fbef0480b575e6dafbaa) chore: bump dependencies
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.4.0](https://github.com/siderolabs/extensions/releases/tag/v1.4.0)

## [Talos System Extensions 1.5.0-alpha.3](https://github.com/siderolabs/extensions/releases/tag/v1.5.0-alpha.3) (2023-07-25)

Welcome to the v1.5.0-alpha.3 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.5/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### QEMU Guest Agent

QEMU Guest Agent is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Tailscale

Tailscale is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Thunderbolt

Thunderbolt drivers is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Component Updates

* Linux firmware: 20230625
* DRBD: 9.2.4
* Intel ucode: 20230613
* Open iSCSI: 2.1.9


### ZFS

ZFS is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Contributors

* Andrey Smirnov
* Noel Georgi
* Markus Reiter
* Roee Klinger
* Andrei Kvapil
* Igor Rzegocki
* beau trepp

### Changes
<details><summary>16 commits</summary>
<p>

* [`1fdca1a`](https://github.com/siderolabs/extensions/commit/1fdca1a892054036c9661351d9740ce9e35eb2b3) feat: add thunderbolt/USB4 module
* [`14c9ec6`](https://github.com/siderolabs/extensions/commit/14c9ec6c9af143f663865d91b9188db0aa3f4e7b) feat: add 'zfs' extension
* [`cca830e`](https://github.com/siderolabs/extensions/commit/cca830ec2f0f29d2a0cf604dd7680bfa9d2a5d69) release(v1.5.0-alpha.2): prepare release
* [`44b53b8`](https://github.com/siderolabs/extensions/commit/44b53b8da53fa25646612ca32944089fc166597e) fix: ci blob upload
* [`6366b17`](https://github.com/siderolabs/extensions/commit/6366b171c75e678a4acf5dd791a4e906ed3ccf04) chore: use machined for shutdown calls
* [`048de33`](https://github.com/siderolabs/extensions/commit/048de338160553383e9e202478281ba14e7355bc) chore: bump dependencies
* [`b98fed7`](https://github.com/siderolabs/extensions/commit/b98fed77b3a9b7891c82a89fe94ce7cd6dc455fc) fix: use correct exit status in `nut-client` patch
* [`1934a5b`](https://github.com/siderolabs/extensions/commit/1934a5b12bfad6915bb8ff7f02c82ce6448c09ce) feat: add `qemu-guest-agent` extension
* [`6c502e1`](https://github.com/siderolabs/extensions/commit/6c502e1095448178b7ff04ebb97cde12fc9b6be8) feat: tailscale extension
* [`30f24da`](https://github.com/siderolabs/extensions/commit/30f24da0ce678b08705fbb9d95b85e44a5f1a353) release(v1.5.0-alpha.1): prepare release
* [`cfee4f6`](https://github.com/siderolabs/extensions/commit/cfee4f6d14c1130496bfdc1baa3c9b4d4958c9bf) chore: bump deps
* [`e814607`](https://github.com/siderolabs/extensions/commit/e8146076896e0b370ce1e1cb02d97d7dd559b0b9) docs: remove the belkin_sa from readme
* [`4cee34c`](https://github.com/siderolabs/extensions/commit/4cee34c5859e17264f899717bf843bbdccff4a4e) feat: add usb modem drivers
* [`f05c997`](https://github.com/siderolabs/extensions/commit/f05c9975c92a15c36f11dce87d05ffa0d7164e10) release(v1.5.0-alpha.0): prepare release
* [`0855dd7`](https://github.com/siderolabs/extensions/commit/0855dd71e7831c51ded46192845a85f4ba019c41) chore: bump dependencies
* [`c06874c`](https://github.com/siderolabs/extensions/commit/c06874c102fe9fc6ceb93a48be77434194f8940c) chore: add make targets to automate image signing
</p>
</details>

### Changes since v1.5.0-alpha.2
<details><summary>2 commits</summary>
<p>

* [`1fdca1a`](https://github.com/siderolabs/extensions/commit/1fdca1a892054036c9661351d9740ce9e35eb2b3) feat: add thunderbolt/USB4 module
* [`14c9ec6`](https://github.com/siderolabs/extensions/commit/14c9ec6c9af143f663865d91b9188db0aa3f4e7b) feat: add 'zfs' extension
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.4.0](https://github.com/siderolabs/extensions/releases/tag/v1.4.0)

## [Talos System Extensions 1.5.0-alpha.2](https://github.com/siderolabs/extensions/releases/tag/v1.5.0-alpha.2) (2023-07-20)

Welcome to the v1.5.0-alpha.2 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.5/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### QEMU Guest Agent

QEMU Guest Agent is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Tailscale

Tailscale is now supported as Talos System Extension. Requires Talos v1.5.0 or later.


### Component Updates

* Linux firmware: 20230625
* DRBD: 9.2.4
* Intel ucode: 20230613
* Open iSCSI: 2.1.9


### Contributors

* Andrey Smirnov
* Noel Georgi
* Markus Reiter
* Roee Klinger
* beau trepp

### Changes
<details><summary>13 commits</summary>
<p>

* [`44b53b8`](https://github.com/siderolabs/extensions/commit/44b53b8da53fa25646612ca32944089fc166597e) fix: ci blob upload
* [`6366b17`](https://github.com/siderolabs/extensions/commit/6366b171c75e678a4acf5dd791a4e906ed3ccf04) chore: use machined for shutdown calls
* [`048de33`](https://github.com/siderolabs/extensions/commit/048de338160553383e9e202478281ba14e7355bc) chore: bump dependencies
* [`b98fed7`](https://github.com/siderolabs/extensions/commit/b98fed77b3a9b7891c82a89fe94ce7cd6dc455fc) fix: use correct exit status in `nut-client` patch
* [`1934a5b`](https://github.com/siderolabs/extensions/commit/1934a5b12bfad6915bb8ff7f02c82ce6448c09ce) feat: add `qemu-guest-agent` extension
* [`6c502e1`](https://github.com/siderolabs/extensions/commit/6c502e1095448178b7ff04ebb97cde12fc9b6be8) feat: tailscale extension
* [`30f24da`](https://github.com/siderolabs/extensions/commit/30f24da0ce678b08705fbb9d95b85e44a5f1a353) release(v1.5.0-alpha.1): prepare release
* [`cfee4f6`](https://github.com/siderolabs/extensions/commit/cfee4f6d14c1130496bfdc1baa3c9b4d4958c9bf) chore: bump deps
* [`e814607`](https://github.com/siderolabs/extensions/commit/e8146076896e0b370ce1e1cb02d97d7dd559b0b9) docs: remove the belkin_sa from readme
* [`4cee34c`](https://github.com/siderolabs/extensions/commit/4cee34c5859e17264f899717bf843bbdccff4a4e) feat: add usb modem drivers
* [`f05c997`](https://github.com/siderolabs/extensions/commit/f05c9975c92a15c36f11dce87d05ffa0d7164e10) release(v1.5.0-alpha.0): prepare release
* [`0855dd7`](https://github.com/siderolabs/extensions/commit/0855dd71e7831c51ded46192845a85f4ba019c41) chore: bump dependencies
* [`c06874c`](https://github.com/siderolabs/extensions/commit/c06874c102fe9fc6ceb93a48be77434194f8940c) chore: add make targets to automate image signing
</p>
</details>

### Changes since v1.5.0-alpha.1
<details><summary>6 commits</summary>
<p>

* [`44b53b8`](https://github.com/siderolabs/extensions/commit/44b53b8da53fa25646612ca32944089fc166597e) fix: ci blob upload
* [`6366b17`](https://github.com/siderolabs/extensions/commit/6366b171c75e678a4acf5dd791a4e906ed3ccf04) chore: use machined for shutdown calls
* [`048de33`](https://github.com/siderolabs/extensions/commit/048de338160553383e9e202478281ba14e7355bc) chore: bump dependencies
* [`b98fed7`](https://github.com/siderolabs/extensions/commit/b98fed77b3a9b7891c82a89fe94ce7cd6dc455fc) fix: use correct exit status in `nut-client` patch
* [`1934a5b`](https://github.com/siderolabs/extensions/commit/1934a5b12bfad6915bb8ff7f02c82ce6448c09ce) feat: add `qemu-guest-agent` extension
* [`6c502e1`](https://github.com/siderolabs/extensions/commit/6c502e1095448178b7ff04ebb97cde12fc9b6be8) feat: tailscale extension
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.4.0](https://github.com/siderolabs/extensions/releases/tag/v1.4.0)

## [Talos System Extensions 1.5.0-alpha.1](https://github.com/siderolabs/extensions/releases/tag/v1.5.0-alpha.1) (2023-06-22)

Welcome to the v1.5.0-alpha.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.5/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

* Linux firmware: 20230515
* DRBD: 9.2.4
* Intel ucode: 20230613


### Contributors

* Andrey Smirnov
* Roee Klinger
* Noel Georgi

### Changes
<details><summary>6 commits</summary>
<p>

* [`cfee4f6`](https://github.com/siderolabs/extensions/commit/cfee4f6d14c1130496bfdc1baa3c9b4d4958c9bf) chore: bump deps
* [`e814607`](https://github.com/siderolabs/extensions/commit/e8146076896e0b370ce1e1cb02d97d7dd559b0b9) docs: remove the belkin_sa from readme
* [`4cee34c`](https://github.com/siderolabs/extensions/commit/4cee34c5859e17264f899717bf843bbdccff4a4e) feat: add usb modem drivers
* [`f05c997`](https://github.com/siderolabs/extensions/commit/f05c9975c92a15c36f11dce87d05ffa0d7164e10) release(v1.5.0-alpha.0): prepare release
* [`0855dd7`](https://github.com/siderolabs/extensions/commit/0855dd71e7831c51ded46192845a85f4ba019c41) chore: bump dependencies
* [`c06874c`](https://github.com/siderolabs/extensions/commit/c06874c102fe9fc6ceb93a48be77434194f8940c) chore: add make targets to automate image signing
</p>
</details>

### Changes since v1.5.0-alpha.0
<details><summary>3 commits</summary>
<p>

* [`cfee4f6`](https://github.com/siderolabs/extensions/commit/cfee4f6d14c1130496bfdc1baa3c9b4d4958c9bf) chore: bump deps
* [`e814607`](https://github.com/siderolabs/extensions/commit/e8146076896e0b370ce1e1cb02d97d7dd559b0b9) docs: remove the belkin_sa from readme
* [`4cee34c`](https://github.com/siderolabs/extensions/commit/4cee34c5859e17264f899717bf843bbdccff4a4e) feat: add usb modem drivers
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.4.0](https://github.com/siderolabs/extensions/releases/tag/v1.4.0)

## [Talos System Extensions 1.5.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.5.0-alpha.0) (2023-05-19)

Welcome to the v1.5.0-alpha.0 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.5/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

* Linux firmware: 20230515
* DRBD: 9.2.3
* Intel ucode: 20230512


### Contributors

* Andrey Smirnov

### Changes
<details><summary>2 commits</summary>
<p>

* [`0855dd7`](https://github.com/siderolabs/extensions/commit/0855dd71e7831c51ded46192845a85f4ba019c41) chore: bump dependencies
* [`c06874c`](https://github.com/siderolabs/extensions/commit/c06874c102fe9fc6ceb93a48be77434194f8940c) chore: add make targets to automate image signing
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.4.0](https://github.com/siderolabs/extensions/releases/tag/v1.4.0)

## [Talos System Extensions 1.4.0-alpha.4](https://github.com/siderolabs/extensions/releases/tag/v1.4.0-alpha.4) (2023-03-31)

Welcome to the v1.4.0-alpha.4 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.4/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

* Gvisor: 20231214.0
* Linux firmware: 20230310
* Intel ucode: 20230214
* nvidia-container-toolkit: v1.12.1
* NVIDIA driver: 530.41.03
* nvidia-fabricmanager: 525.85.12


### Contributors

* Noel Georgi
* Andrey Smirnov
* Brandon McNama
* Jori Huisman
* Spencer Smith
* Utku Ozdemir

### Changes
<details><summary>24 commits</summary>
<p>

* [`3874e3d`](https://github.com/siderolabs/extensions/commit/3874e3d2bf1e53b972306ddf1aa756a25624c90f) chore: bump deps
* [`e3e4012`](https://github.com/siderolabs/extensions/commit/e3e4012ea21957e323fc1e848211bedcbe3055a9) fix: nvidia modules autoload
* [`666e073`](https://github.com/siderolabs/extensions/commit/666e073929aaa8f1da756d081b9301c5f7fb1ba8) release(v1.4.0-alpha.3): prepare release
* [`130ebd5`](https://github.com/siderolabs/extensions/commit/130ebd5798af180908e501b4ead3494421a43026) chore: bump deps
* [`1ceb958`](https://github.com/siderolabs/extensions/commit/1ceb95857579157a13589f396bc4683f05d55065) docs: fix minor typo
* [`75cf535`](https://github.com/siderolabs/extensions/commit/75cf5350ffc8b08f1167e291b8d7e93295c2ff38) feat: update Go 1.20.2
* [`58cf294`](https://github.com/siderolabs/extensions/commit/58cf29467048e56fd71fc8fa6647263e3c75c0e6) chore: use azure blob storage for builds
* [`29859fa`](https://github.com/siderolabs/extensions/commit/29859fad8e81da83b13e951b0d1b67450ef26883) release(v1.4.0-alpha.2): prepare release
* [`be10780`](https://github.com/siderolabs/extensions/commit/be107808f4c7cff8ef4a986081433f3228a4d6d0) chore: re-enable gvisor cgroup support
* [`7cdf5e7`](https://github.com/siderolabs/extensions/commit/7cdf5e7df4519c6a2010082e5a2fd7bde5ee6b8d) chore: use standard `PKGS` variable
* [`1f10906`](https://github.com/siderolabs/extensions/commit/1f1090606f030b6615d78dfa6ed6e54a369486b9) fix: drone artifact path
* [`53db8a8`](https://github.com/siderolabs/extensions/commit/53db8a825156a903862bb0287bb3eb5b73a594cf) feat: add drone promote e2e stage
* [`85e727c`](https://github.com/siderolabs/extensions/commit/85e727c172e9b5303d231af2b24756eb73da7fcf) chore: make `PKGS_VERSION` a variable
* [`370a093`](https://github.com/siderolabs/extensions/commit/370a093e2f41477d11e2b7fe7219f17b6be5cb8d) fix: renovate regex
* [`8cb8014`](https://github.com/siderolabs/extensions/commit/8cb8014ce27bf3fb3b3cde89f419241d468d5396) chore: bump deps
* [`1d9465c`](https://github.com/siderolabs/extensions/commit/1d9465c5b8c5e3cb540c34938b0e1543cd9b5a34) feat: add i915 microcode
* [`b378a46`](https://github.com/siderolabs/extensions/commit/b378a465fe7d3a78c068b20f97a2f5e43e6227e2) feat: add mellanox ofed drivers extension
* [`2126f28`](https://github.com/siderolabs/extensions/commit/2126f28338b63442a8ced83d388cce7bdcbda3bb) release(v1.4.0-alpha.1): prepare release
* [`f5eefa7`](https://github.com/siderolabs/extensions/commit/f5eefa721ebf1d7d945f1bcd39a072e6ad56b32b) chore: bump dependencies
* [`8a55387`](https://github.com/siderolabs/extensions/commit/8a5538704273ce02ae4bb233d7db3d54feda17f3) chore: disable provenance in buildx
* [`bcbbf43`](https://github.com/siderolabs/extensions/commit/bcbbf4337b07fafce73c56bd98a2458252fcce45) chore: use default symlinks from base
* [`932a49e`](https://github.com/siderolabs/extensions/commit/932a49e12072ece4966ffff066f72d8c0f43b289) feat: update releases
* [`2261aa4`](https://github.com/siderolabs/extensions/commit/2261aa42065380d1a25000fa2d7d201fe035baa0) fix: iscsi-tools paths
* [`a72fd6c`](https://github.com/siderolabs/extensions/commit/a72fd6c7804dd3ff5e97bf78068ab02b2346efb5) fix: iscsi-tools extension
</p>
</details>

### Changes since v1.4.0-alpha.3
<details><summary>2 commits</summary>
<p>

* [`3874e3d`](https://github.com/siderolabs/extensions/commit/3874e3d2bf1e53b972306ddf1aa756a25624c90f) chore: bump deps
* [`e3e4012`](https://github.com/siderolabs/extensions/commit/e3e4012ea21957e323fc1e848211bedcbe3055a9) fix: nvidia modules autoload
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.3.6](https://github.com/siderolabs/extensions/releases/tag/v1.3.6)

## [Talos System Extensions 1.4.0-alpha.3](https://github.com/siderolabs/extensions/releases/tag/v1.4.0-alpha.3) (2023-03-23)

Welcome to the v1.4.0-alpha.3 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.4/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

* Gvisor: 20231214.0
* Linux firmware: 20230310
* Intel ucode: 20230214
* nvidia-container-toolkit: v1.12.1
* NVIDIA driver: 530.30.02
* nvidia-fabricmanager: 525.85.12


### Contributors

* Noel Georgi
* Andrey Smirnov
* Brandon McNama
* Jori Huisman
* Spencer Smith
* Utku Ozdemir

### Changes
<details><summary>21 commits</summary>
<p>

* [`130ebd5`](https://github.com/siderolabs/extensions/commit/130ebd5798af180908e501b4ead3494421a43026) chore: bump deps
* [`1ceb958`](https://github.com/siderolabs/extensions/commit/1ceb95857579157a13589f396bc4683f05d55065) docs: fix minor typo
* [`75cf535`](https://github.com/siderolabs/extensions/commit/75cf5350ffc8b08f1167e291b8d7e93295c2ff38) feat: update Go 1.20.2
* [`58cf294`](https://github.com/siderolabs/extensions/commit/58cf29467048e56fd71fc8fa6647263e3c75c0e6) chore: use azure blob storage for builds
* [`29859fa`](https://github.com/siderolabs/extensions/commit/29859fad8e81da83b13e951b0d1b67450ef26883) release(v1.4.0-alpha.2): prepare release
* [`be10780`](https://github.com/siderolabs/extensions/commit/be107808f4c7cff8ef4a986081433f3228a4d6d0) chore: re-enable gvisor cgroup support
* [`7cdf5e7`](https://github.com/siderolabs/extensions/commit/7cdf5e7df4519c6a2010082e5a2fd7bde5ee6b8d) chore: use standard `PKGS` variable
* [`1f10906`](https://github.com/siderolabs/extensions/commit/1f1090606f030b6615d78dfa6ed6e54a369486b9) fix: drone artifact path
* [`53db8a8`](https://github.com/siderolabs/extensions/commit/53db8a825156a903862bb0287bb3eb5b73a594cf) feat: add drone promote e2e stage
* [`85e727c`](https://github.com/siderolabs/extensions/commit/85e727c172e9b5303d231af2b24756eb73da7fcf) chore: make `PKGS_VERSION` a variable
* [`370a093`](https://github.com/siderolabs/extensions/commit/370a093e2f41477d11e2b7fe7219f17b6be5cb8d) fix: renovate regex
* [`8cb8014`](https://github.com/siderolabs/extensions/commit/8cb8014ce27bf3fb3b3cde89f419241d468d5396) chore: bump deps
* [`1d9465c`](https://github.com/siderolabs/extensions/commit/1d9465c5b8c5e3cb540c34938b0e1543cd9b5a34) feat: add i915 microcode
* [`b378a46`](https://github.com/siderolabs/extensions/commit/b378a465fe7d3a78c068b20f97a2f5e43e6227e2) feat: add mellanox ofed drivers extension
* [`2126f28`](https://github.com/siderolabs/extensions/commit/2126f28338b63442a8ced83d388cce7bdcbda3bb) release(v1.4.0-alpha.1): prepare release
* [`f5eefa7`](https://github.com/siderolabs/extensions/commit/f5eefa721ebf1d7d945f1bcd39a072e6ad56b32b) chore: bump dependencies
* [`8a55387`](https://github.com/siderolabs/extensions/commit/8a5538704273ce02ae4bb233d7db3d54feda17f3) chore: disable provenance in buildx
* [`bcbbf43`](https://github.com/siderolabs/extensions/commit/bcbbf4337b07fafce73c56bd98a2458252fcce45) chore: use default symlinks from base
* [`932a49e`](https://github.com/siderolabs/extensions/commit/932a49e12072ece4966ffff066f72d8c0f43b289) feat: update releases
* [`2261aa4`](https://github.com/siderolabs/extensions/commit/2261aa42065380d1a25000fa2d7d201fe035baa0) fix: iscsi-tools paths
* [`a72fd6c`](https://github.com/siderolabs/extensions/commit/a72fd6c7804dd3ff5e97bf78068ab02b2346efb5) fix: iscsi-tools extension
</p>
</details>

### Changes since v1.4.0-alpha.2
<details><summary>4 commits</summary>
<p>

* [`130ebd5`](https://github.com/siderolabs/extensions/commit/130ebd5798af180908e501b4ead3494421a43026) chore: bump deps
* [`1ceb958`](https://github.com/siderolabs/extensions/commit/1ceb95857579157a13589f396bc4683f05d55065) docs: fix minor typo
* [`75cf535`](https://github.com/siderolabs/extensions/commit/75cf5350ffc8b08f1167e291b8d7e93295c2ff38) feat: update Go 1.20.2
* [`58cf294`](https://github.com/siderolabs/extensions/commit/58cf29467048e56fd71fc8fa6647263e3c75c0e6) chore: use azure blob storage for builds
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.3.0](https://github.com/siderolabs/extensions/releases/tag/v1.3.0)

## [Talos System Extensions 1.4.0-alpha.2](https://github.com/siderolabs/extensions/releases/tag/v1.4.0-alpha.2) (2023-02-28)

Welcome to the v1.4.0-alpha.2 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.2/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

* Gvisor: 20231214.0
* Linux firmware: 20230210
* Intel ucode: 20230214
* nvidia-container-toolkit: v1.12.0
* nvidia driver: 525.89.02
* nvidia-fabricmanager: 525.85.12


### Contributors

* Noel Georgi
* Andrey Smirnov
* Brandon McNama
* Utku Ozdemir

### Changes
<details><summary>16 commits</summary>
<p>

* [`be10780`](https://github.com/siderolabs/extensions/commit/be107808f4c7cff8ef4a986081433f3228a4d6d0) chore: re-enable gvisor cgroup support
* [`7cdf5e7`](https://github.com/siderolabs/extensions/commit/7cdf5e7df4519c6a2010082e5a2fd7bde5ee6b8d) chore: use standard `PKGS` variable
* [`1f10906`](https://github.com/siderolabs/extensions/commit/1f1090606f030b6615d78dfa6ed6e54a369486b9) fix: drone artifact path
* [`53db8a8`](https://github.com/siderolabs/extensions/commit/53db8a825156a903862bb0287bb3eb5b73a594cf) feat: add drone promote e2e stage
* [`85e727c`](https://github.com/siderolabs/extensions/commit/85e727c172e9b5303d231af2b24756eb73da7fcf) chore: make `PKGS_VERSION` a variable
* [`370a093`](https://github.com/siderolabs/extensions/commit/370a093e2f41477d11e2b7fe7219f17b6be5cb8d) fix: renovate regex
* [`8cb8014`](https://github.com/siderolabs/extensions/commit/8cb8014ce27bf3fb3b3cde89f419241d468d5396) chore: bump deps
* [`1d9465c`](https://github.com/siderolabs/extensions/commit/1d9465c5b8c5e3cb540c34938b0e1543cd9b5a34) feat: add i915 microcode
* [`b378a46`](https://github.com/siderolabs/extensions/commit/b378a465fe7d3a78c068b20f97a2f5e43e6227e2) feat: add mellanox ofed drivers extension
* [`2126f28`](https://github.com/siderolabs/extensions/commit/2126f28338b63442a8ced83d388cce7bdcbda3bb) release(v1.4.0-alpha.1): prepare release
* [`f5eefa7`](https://github.com/siderolabs/extensions/commit/f5eefa721ebf1d7d945f1bcd39a072e6ad56b32b) chore: bump dependencies
* [`8a55387`](https://github.com/siderolabs/extensions/commit/8a5538704273ce02ae4bb233d7db3d54feda17f3) chore: disable provenance in buildx
* [`bcbbf43`](https://github.com/siderolabs/extensions/commit/bcbbf4337b07fafce73c56bd98a2458252fcce45) chore: use default symlinks from base
* [`932a49e`](https://github.com/siderolabs/extensions/commit/932a49e12072ece4966ffff066f72d8c0f43b289) feat: update releases
* [`2261aa4`](https://github.com/siderolabs/extensions/commit/2261aa42065380d1a25000fa2d7d201fe035baa0) fix: iscsi-tools paths
* [`a72fd6c`](https://github.com/siderolabs/extensions/commit/a72fd6c7804dd3ff5e97bf78068ab02b2346efb5) fix: iscsi-tools extension
</p>
</details>

### Changes since v1.4.0-alpha.1
<details><summary>9 commits</summary>
<p>

* [`be10780`](https://github.com/siderolabs/extensions/commit/be107808f4c7cff8ef4a986081433f3228a4d6d0) chore: re-enable gvisor cgroup support
* [`7cdf5e7`](https://github.com/siderolabs/extensions/commit/7cdf5e7df4519c6a2010082e5a2fd7bde5ee6b8d) chore: use standard `PKGS` variable
* [`1f10906`](https://github.com/siderolabs/extensions/commit/1f1090606f030b6615d78dfa6ed6e54a369486b9) fix: drone artifact path
* [`53db8a8`](https://github.com/siderolabs/extensions/commit/53db8a825156a903862bb0287bb3eb5b73a594cf) feat: add drone promote e2e stage
* [`85e727c`](https://github.com/siderolabs/extensions/commit/85e727c172e9b5303d231af2b24756eb73da7fcf) chore: make `PKGS_VERSION` a variable
* [`370a093`](https://github.com/siderolabs/extensions/commit/370a093e2f41477d11e2b7fe7219f17b6be5cb8d) fix: renovate regex
* [`8cb8014`](https://github.com/siderolabs/extensions/commit/8cb8014ce27bf3fb3b3cde89f419241d468d5396) chore: bump deps
* [`1d9465c`](https://github.com/siderolabs/extensions/commit/1d9465c5b8c5e3cb540c34938b0e1543cd9b5a34) feat: add i915 microcode
* [`b378a46`](https://github.com/siderolabs/extensions/commit/b378a465fe7d3a78c068b20f97a2f5e43e6227e2) feat: add mellanox ofed drivers extension
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.3.0](https://github.com/siderolabs/extensions/releases/tag/v1.3.0)

## [Talos System Extensions 1.4.0-alpha.1](https://github.com/siderolabs/extensions/releases/tag/v1.4.0-alpha.1) (2023-01-26)

Welcome to the v1.4.0-alpha.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.2/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Component Updates

* Gvisor: 20230109.0
* Linux firmware: 20221214


### Contributors

* Noel Georgi
* Andrey Smirnov
* Utku Ozdemir

### Changes
<details><summary>6 commits</summary>
<p>

* [`f5eefa7`](https://github.com/siderolabs/extensions/commit/f5eefa721ebf1d7d945f1bcd39a072e6ad56b32b) chore: bump dependencies
* [`8a55387`](https://github.com/siderolabs/extensions/commit/8a5538704273ce02ae4bb233d7db3d54feda17f3) chore: disable provenance in buildx
* [`bcbbf43`](https://github.com/siderolabs/extensions/commit/bcbbf4337b07fafce73c56bd98a2458252fcce45) chore: use default symlinks from base
* [`932a49e`](https://github.com/siderolabs/extensions/commit/932a49e12072ece4966ffff066f72d8c0f43b289) feat: update releases
* [`2261aa4`](https://github.com/siderolabs/extensions/commit/2261aa42065380d1a25000fa2d7d201fe035baa0) fix: iscsi-tools paths
* [`a72fd6c`](https://github.com/siderolabs/extensions/commit/a72fd6c7804dd3ff5e97bf78068ab02b2346efb5) fix: iscsi-tools extension
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.3.0](https://github.com/siderolabs/extensions/releases/tag/v1.3.0)

## [Talos System Extensions 1.3.0-alpha.1](https://github.com/siderolabs/extensions/releases/tag/v1.3.0-alpha.1) (2022-10-31)

Welcome to the v1.3.0-alpha.1 release of Talos System Extensions!  
*This is a pre-release of Talos System Extensions*

See [Talos Linux documentation](https://www.talos.dev/v1.2/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Linbit DRDB driver

Extensions now publish the [DRBD driver](https://linbit.com/drbd/)


### Coral.ai TPU gasket driver

Extensions now publish the gasket driver for [Coral.ai](https://coral.ai/) PCIe TPU modules.


### Component Updates

* Gvisor: 20220919.0
* Linux firmware: 20220913
* NVIDIA container-toolkit: v1.11.0
* NVIDIA libnvidia-container: v1.11.0
* nvidia-container-toolkit system extension: v1.11.0
* Open iSCSI: 2.1.7
* Open iSNS: 0.102
* iscsi-tools system extension: v0.1.2


### Contributors

* Andrey Smirnov
* DJAlPee
* Tim Jones

### Changes
<details><summary>3 commits</summary>
<p>

* [`0ea902e`](https://github.com/siderolabs/extensions/commit/0ea902e326091246e38b490660a15850cec7d933) chore: update dependencies
* [`a9f42bb`](https://github.com/siderolabs/extensions/commit/a9f42bb207df64e9ed598c5394c19ca9ff9b3e1a) feat: add 'drbd' extension
* [`eff6a5d`](https://github.com/siderolabs/extensions/commit/eff6a5d048cb50984492b620864c4e669bc4faa5) feat: update releases
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.3.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.3.0-alpha.0)

## [Talos System Extensions 1.3.0-alpha.0](https://github.com/siderolabs/extensions/releases/tag/v1.3.0-alpha.0) (2022-09-29)

Welcome to the v1.3.0-alpha.0 release of Talos System Extensions!

See [Talos Linux documentation](https://www.talos.dev/v1.2/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Coral.ai TPU gasket driver

Extensions now publish the gasket driver for [Coral.ai](https://coral.ai/) PCIe TPU modules.


### Component Updates

* Gvisor: 20220919.0
* Linux firmware: 20220913
* NVIDIA container-toolkit: v1.11.0
* NVIDIA libnvidia-container: v1.11.0
* nvidia-container-toolkit system extension: v1.11.0
* Open iSCSI: 2.1.7
* Open iSNS: 0.102
* iscsi-tools system extension: v0.1.2


### Contributors

* Branden Cash
* Noel Georgi

### Changes
<details><summary>2 commits</summary>
<p>

* [`d428c1f`](https://github.com/siderolabs/extensions/commit/d428c1f81d3252793c1c3c5fa047bc92f68079f0) feat: add gasket driver
* [`b4edb73`](https://github.com/siderolabs/extensions/commit/b4edb73cd4dd1f2cc06f9357ae67305213ae2537) chore: bump deps
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.2.3](https://github.com/siderolabs/extensions/releases/tag/v1.2.3)

## [Talos System Extensions 1.2.3](https://github.com/siderolabs/extensions/releases/tag/v1.2.3) (2022-09-20)

Welcome to the v1.2.3 release of Talos System Extensions!

See [Talos Linux documentation](https://www.talos.dev/v1.2/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Contributors

* Noel Georgi

### Changes
<details><summary>1 commit</summary>
<p>

* [`0991e6c`](https://github.com/siderolabs/extensions/commit/0991e6c8a06f158ada10ae98150f15801ff755b3) fix: drone release pipeline
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.2.2](https://github.com/siderolabs/extensions/releases/tag/v1.2.2)

## [Talos System Extensions 1.2.2](https://github.com/siderolabs/extensions/releases/tag/v1.2.2) (2022-09-13)

Welcome to the v1.2.2 release of Talos System Extensions!

See [Talos Linux documentation](https://www.talos.dev/v1.2/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Power management

Support for Network UPS Tools client system extension to handle node graceful shutdown on power events from supported UPS's.


### Contributors

* Daniel Quinlan

### Changes
<details><summary>1 commit</summary>
<p>

* [`47bd65c`](https://github.com/siderolabs/extensions/commit/47bd65c61bb38b1512c6b90d4b704622fad4b588) feat: add nut-client
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.2.1](https://github.com/siderolabs/extensions/releases/tag/v1.2.1)

## [Talos System Extensions 1.2.1](https://github.com/siderolabs/extensions/releases/tag/v1.2.1) (2022-09-07)

Welcome to the 1.2.1 release of Talos System Extensions!

See [Talos Linux documentation](https://www.talos.dev/v1.2/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Updates

Update to support Talos v1.2.1


### Contributors

* Noel Georgi

### Changes
<details><summary>1 commit</summary>
<p>

* [`b5be6a9`](https://github.com/siderolabs/extensions/commit/b5be6a9ce178b6ef7d9f3653a9d9e2abfe5b29e5) chore: fix changelog
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.2.0](https://github.com/siderolabs/extensions/releases/tag/v1.2.0)

## [Talos System Extensions 1.2.0](https://github.com/siderolabs/extensions/releases/tag/v1.2.0) (2022-09-01)

Welcome to the 1.2.0 release of Talos System Extensions!

See [Talos Linux documentation](https://www.talos.dev/v1.0/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Talos Linux Compatibility

System extensions are compatible with Talos v1.0.x.


### Contributors

* Noel Georgi
* Andrey Smirnov
* Andrew Rynhard
* Philipp Sauter

### Changes
<details><summary>23 commits</summary>
<p>

* [`051dc04`](https://github.com/siderolabs/extensions/commit/051dc043ee2d91dc94a3ebf53ff36696a97bc98c) release(v1.2.0): prepare release
* [`12070a7`](https://github.com/siderolabs/extensions/commit/12070a7f018e56691d34171dcf2df87e01f75e93) release(v1.2.0-beta.2): prepare release
* [`186d372`](https://github.com/siderolabs/extensions/commit/186d37219835eeb35597700b89cebe488a54139a) chore: fix regex for `vars.yaml`
* [`eac3211`](https://github.com/siderolabs/extensions/commit/eac3211468a2f9cac7f3d967fb8507e39ba81e13) feat: enable renovate bot
* [`036b63f`](https://github.com/siderolabs/extensions/commit/036b63f4bb24aa767414e6d97c6a2ba72e65e8b1) chore: bump intel ucode to 20220809
* [`0a5d92b`](https://github.com/siderolabs/extensions/commit/0a5d92bb7f923849200525b95577295a8b362839) fix: nvidia pkg versioning
* [`e77f347`](https://github.com/siderolabs/extensions/commit/e77f3477ee5633f259a71504ee514ed240b472c8) feat: publish nvidia modules and toolkit
* [`adbfba9`](https://github.com/siderolabs/extensions/commit/adbfba979118e04b9eb70d315d18155fe1326569) feat: allow modules to be loaded via extension
* [`93a7a0d`](https://github.com/siderolabs/extensions/commit/93a7a0df54362a4279eb5cf9fc83c1f3023e9648) chore: bump nvidia drivers to 515.65.01
* [`fece6ed`](https://github.com/siderolabs/extensions/commit/fece6ed56fc8119dc672088de69533c1137900b3) chore: update extension spec
* [`6c412b0`](https://github.com/siderolabs/extensions/commit/6c412b03c4435014e3e85935e9ba3855ec7ddaed) fix: parsing pidfile with newlines
* [`14946ee`](https://github.com/siderolabs/extensions/commit/14946ee3afe3090bd8f0da80125f05f8d243eaf2) chore: update Go to 1.18.4
* [`b5de7b4`](https://github.com/siderolabs/extensions/commit/b5de7b49e8b39a9ade689b0ed89fda592f86ab2c) chore: bump nvidia drivers to 515.43.04
* [`82b41ad`](https://github.com/siderolabs/extensions/commit/82b41adcff44310f3832efd59ea57767750ece4a) chore: bump golang to 1.18.3
* [`925af3b`](https://github.com/siderolabs/extensions/commit/925af3b45c7cf2c6131dc8187075a5c37e33e8b3) chore: update pkgs to 1.1.0
* [`ef860ab`](https://github.com/siderolabs/extensions/commit/ef860abc71665385467e9c63f0076b2ec44ce9c0) chore: stable source date epoch
* [`e0b1465`](https://github.com/siderolabs/extensions/commit/e0b1465c5ce60aa6c761b2a75ea54a8b69f130f6) chore: add iscsi extra perms
* [`7cf2843`](https://github.com/siderolabs/extensions/commit/7cf2843cd73aa807d72151ec64bb0dc913492d63) chore: nvidia-persistenced as an extension service
* [`3ccc1b5`](https://github.com/siderolabs/extensions/commit/3ccc1b54a919ac32c26195df26afe6dc0372dbca) fix: copy the hello world extension spec to the correct place
* [`7c92e63`](https://github.com/siderolabs/extensions/commit/7c92e637ff6c6ab49ba851ba6c6859817107b737) feat: update Interl microcode to 20220510
* [`00a8b53`](https://github.com/siderolabs/extensions/commit/00a8b53b1fdad57a670896d71e35f0221ebd6e43) chore: use golang.org/x/sys over syscall
* [`dead35b`](https://github.com/siderolabs/extensions/commit/dead35b5b8714a46a1cc999472fc9db76c79d071) chore: move iscsi extension into storage category
* [`182c3d2`](https://github.com/siderolabs/extensions/commit/182c3d2cf7f5bd0c915e39b3e8f3c78024f06410) feat: add open-iscsi system extension
</p>
</details>

### Dependency Changes

This release has no dependency changes

Previous release can be found at [v1.0.0](https://github.com/siderolabs/extensions/releases/tag/v1.0.0)


## [Talos System Extensions 1.0.0](https://github.com/siderolabs/extensions/releases/tag/v1.0.0) (2022-04-25)

Welcome to the v1.0.0 release of Talos System Extensions!

See [Talos Linux documentation](https://www.talos.dev/v1.0/talos-guides/configuration/system-extensions/) for information on using system extensions.

Please try out the release binaries and report any issues at
https://github.com/siderolabs/extensions/issues.

### Talos Linux Compatibility

System extensions are compatible with Talos v1.0.x.


### Contributors

* Andrey Smirnov
* Noel Georgi
* Spencer Smith
* Andrew Rynhard
* Andrew Rynhard

### Changes
<details><summary>23 commits</summary>
<p>

* [`b09dcb8`](https://github.com/siderolabs/extensions/commit/b09dcb832355f0103fccff9c04a315dd77c86328) docs: provide a minimal extension catalog
* [`4b5bac7`](https://github.com/siderolabs/extensions/commit/4b5bac703665aaea9f5c7cdddb2b1e2e91e7e6c0) feat: update Intel microcode to 20220419
* [`cdc0dcf`](https://github.com/siderolabs/extensions/commit/cdc0dcf7fb6512a279e1e6080dff4f6796a8045b) chore: bump linux-firmware to 20220411
* [`ed63195`](https://github.com/siderolabs/extensions/commit/ed63195910d1fba1bc534d85693768e660c982e6) refactor: clean up extensions versioning
* [`5b1a5d7`](https://github.com/siderolabs/extensions/commit/5b1a5d75735d0b7e6bda07e306e2f6348e52fe37) refactor: use `base` image, bump bldr, update versions
* [`3182ab8`](https://github.com/siderolabs/extensions/commit/3182ab8dd72a99b69429fa69eb5f2e837cd38cfa) chore: bump nvidia drivers to 510.60.02
* [`215aa82`](https://github.com/siderolabs/extensions/commit/215aa82134f8dcd2ae542586585a80c20185b76d) feat: add nvidia-container-runtime
* [`495cabb`](https://github.com/siderolabs/extensions/commit/495cabbf75760a58295e11d24aa17684296ae451) feat: allow hardlinks in the rootfs
* [`c2ace06`](https://github.com/siderolabs/extensions/commit/c2ace0627bb0eaa84344172461020c1445ceacd4) feat: update spec to relax service restrictions
* [`80f8f04`](https://github.com/siderolabs/extensions/commit/80f8f041ab800e004930a64a0268fff7cf6742c6) chore: bump tools
* [`a05f558`](https://github.com/siderolabs/extensions/commit/a05f558f20bc42d38fe0b86c6e4ee6edce80d5a1) feat: add example extension hello world service
* [`0757dc6`](https://github.com/siderolabs/extensions/commit/0757dc672b9bbdf2c6b7507605a3e4f19a2948a6) chore: bump golang from tools to 1.17.7
* [`119fab1`](https://github.com/siderolabs/extensions/commit/119fab1ddd40da9770fed93f4ac1cd628c284b05) chore: provider better folder structure
* [`049b9f9`](https://github.com/siderolabs/extensions/commit/049b9f96cd260d5cd00956e0498a7b714ce42ee0) feat: update intel microcode to 20220207
* [`33f613e`](https://github.com/siderolabs/extensions/commit/33f613e69158155b3065bfc2ee94d1109277d188) feat: add broadcom netextreme 2 (bnx2/bnx2x) extension
* [`d4628b6`](https://github.com/siderolabs/extensions/commit/d4628b6e1d6a9ad3acbb28eebd0584eb74cc239e) feat: add amd-ucode extension
* [`54b831d`](https://github.com/siderolabs/extensions/commit/54b831de7ee9f75ac53962fce95514190dddd6ad) chore: fix version requirement for intel-ucode
* [`933cdb8`](https://github.com/siderolabs/extensions/commit/933cdb8c0a7e0766c8bae9f7cad7a4d89cf163a3) docs: update extension spec
* [`33082a2`](https://github.com/siderolabs/extensions/commit/33082a29bac46ef43d5b280363f02cc2fb9c775a) feat: update gvisor extension after testing
* [`c927b54`](https://github.com/siderolabs/extensions/commit/c927b544aaad0e16c7adad0f44417109cc1326f6) feat: add extension for intel-ucode
* [`97eb468`](https://github.com/siderolabs/extensions/commit/97eb4683a69e22acfaa3a0284fe0c4b11932a4ce) feat: add README
* [`81b2fd3`](https://github.com/siderolabs/extensions/commit/81b2fd3cc324ee66d9c7666fa473c60ce3642614) feat: add gvisor
* [`a46b3f2`](https://github.com/siderolabs/extensions/commit/a46b3f24d158614d582da5e6e7e34b596d10cb8e) Initial commit
</p>
</details>

### Dependency Changes

This release has no dependency changes

