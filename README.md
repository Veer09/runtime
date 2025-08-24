# runtime

**runtime** is a minimal, educational container runtime built from scratch to explore how Linux containers work under the hood. Inspired by tools like `runc`, this runtime focuses on learning and implementing key concepts such as namespaces, cgroups, user isolation, filesystem mounting, and more.

## 📌 TODO — Make Runtime Fully OCI-Compatible

### 📦 OCI Spec Compatibility
- [ ] Parse and validate `config.json` as per OCI runtime-spec
- [ ] Support full OCI lifecycle commands:
  - [ ] `create`
  - [ ] `start`
  - [ ] `kill`
  - [ ] `delete`
  - [ ] `state`
  - [ ] `run`
  - [ ] `exec`
- [ ] Generate valid `state.json` output for `state` command

### 🗂 Filesystem & Mounts
- [ ] Mount filesystems and bind mounts as defined in `config.json`
- [ ] Handle `readonly` and `mountPropagation` options
- [ ] Mount `/proc`, `/sys`, and other required virtual filesystems
- [ ] Set correct `rootfs` based on OCI bundle layout

### ⚙️ Namespaces & Process Execution
- [ ] Create and configure Linux namespaces (UTS, PID, Mount, IPC, Network, User)
- [ ] Support joining existing namespaces via `config.json`
- [ ] Map UID/GID for user namespace
- [ ] Set hostname, working directory, env vars, and args from spec

### 🧱 Cgroups & Resource Limits
- [ ] Create and apply cgroup v2 hierarchy
- [ ] Enforce CPU and memory limits
- [ ] Handle `oomScoreAdj`, `pids.limit`, and other optional resource settings

### 🔐 Security & Isolation
- [ ] Drop Linux capabilities according to spec
- [ ] Apply seccomp filters from `config.json`
- [ ] Support AppArmor and SELinux profiles if defined
- [ ] Set `no_new_privileges` if specified

### 🌐 Networking
- [ ] Configure network namespace (if defined)
- [ ] Set up loopback and optional interfaces
- [ ] (Optional) Bridge + veth support for external networking

### ⚙️ Hooks & Lifecycle
- [ ] Execute lifecycle hooks (`prestart`, `poststart`, `poststop`) as defined in `config.json`
- [ ] Pass correct state data to hooks via stdin

### 🧪 Validation & Compliance
- [ ] Pass all [OCI runtime validation tests](https://github.com/opencontainers/runtime-tools)
- [ ] Handle all required error codes and exit statuses
- [ ] Print required OCI-compliant logs and errors




