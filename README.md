<div align="center">
  <img src="./.github/assets/Ephemera.png" alt="Ephemera logo" width="300">
</div>

<h1></h1>

<div align="center">

![Commit](https://img.shields.io/github/last-commit/MounTemed/Ephemera?style=for-the-badge&logo=git&logoColor=D9E0EE&labelColor=0d1117&color=2b3946)
![Tag](https://img.shields.io/github/v/tag/MounTemed/Ephemera?style=for-the-badge&logo=github&logoColor=D9E0EE&labelColor=0d1117&color=22241c)
![Folder](https://img.shields.io/github/languages/code-size/MounTemed/Ephemera?style=for-the-badge&logo=protondrive&logoColor=D9E0EE&labelColor=0d1117&color=2b3946)
![Stars](https://img.shields.io/github/stars/MounTemed/Ephemera?style=for-the-badge&logo=andela&logoColor=D9E0EE&labelColor=0d1117&color=22241c)

<p>A complex CI/CD process for creating a highly reliable application</p>

</div>

<h1></h1>

<div align="center">
  <h3> ❄️ Overview ✨ </h3>
</div>

<details>
<summary>✨ Features</summary>

- **Effortless Installation** – Get a fully functional Kubernetes orchestrator up and running quickly and smoothly
- **Perfect Reproducibility** – Simply replace the secrets, and you're guaranteed a 100% idempotent, working environment every single time. Consistency is key
- **Rigorous CI/CD pipeline** - for every pull request, we run the most thorough and demanding tests to ensure unrivaled reliability

</details>

<details>
<summary>🛠️ Technologies</summary>

| Technology                        | Purpose                                                |
|-----------------------------------|--------------------------------------------------------|
| **Golang**                        | Foundation for high-performance application code       |
| **Podman**                        | Builds secure, rootless containers                     |
| **CRI-O**                         | Lightweight, Kubernetes-native container runtime       |
| **Kubernetes**                    | For powerful orchestration of containerized workloads  |
| **Helm**                          | Simplifies deployment with handy charts                |
| **Ansible**                       | Automates setup for simple, automated deployment       |
| **Nexus Repository**              | Securely stores and manages container images           |
| **ArgoCD**                        | GitOps for automated application delivery              |
| **Cilium**                        | Provides networking and security with eBPF             |
| **Grafana**                       | Visualizes metrics on customizable dashboards          |
| **Prometheus & node-exporter**    | Comprehensive system and application monitoring        |
| **Alloy**                         | Optimized data collection                              |
| **Traefik**                       | Fast reverse proxy server                              |
| **GitLab CE**                     | Advanced, self-hosted CI/CD pipeline                   |

</details>

<details>
<summary>📦 Installation</summary>

1. **Copy the repository**:

  ```bash
  git clone https://github.com/MounTemed/ephemera
  ```


2. **Go to the downloaded repository**:

  ```bash
  cd ephemera
  ```


3. **Configure infra/ansible/inventory/hosts.yaml to the desired IPv4**:

  ```yaml
  russian:
    children:
      peaceful:
        children:
          prod:
            hosts:
              host1:
                ansible_host: 192.168.0.1 # There should be something in production, but we don't have production
                ansible_user: root
          test:
            hosts:
              host1:
                ansible_host: 192.168.0.2 # CHANGE
                ansible_user: root
          dev:
            hosts:
              host1:
                ansible_host: 192.168.0.3 # CHANGE
                ansible_user: root
  ```

4. **You can run**:

  ```bash
  ansible-playbook infra/ansible/playbook/master.yaml --limit dev
  ```

</details>

<h1></h1>

<div align="center">
  <h3> ✨ Screenshots ❄️ </h3>
</div>

<div align="center">
  <p>Full control, do whatever you want with monitoring 💫</p>
  <img src="./.github/assets/grafana.png">
</div>

<div align="center">
  <p>Just one commit and the code is already on the server! Update without problems, SLO 99% ❄️</p>
  <img src="./.github/assets/argo.png">
</div>

<div align="center">
  <p>GitLab CI is extremely simple, yet very functional ✨</p>
  <img src="./.github/assets/gitlab.png">
</div>
