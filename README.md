# Kubernetes Inventory Taker (AMOS SS 2023)

## Project Mission

Simplifying Kubernetes Management and Monitoring for DevOps Teams

At KIT (Kubernetes Inventory Taker), our mission is to simplify Kubernetes management and monitoring for DevOps teams. We want to provide a user-friendly tool that gives you a clear, real-time view of your Kubernetes resources, from deployments and pods to containers, so you can easily manage and optimize your infrastructure.

Here's what we're all about:

- Inventory Visibility: We'll provide you with a comprehensive view of your Kubernetes inventory, showing you the state, health, and configuration of your deployments, pods, and containers in a single, easy-to-navigate interface. No more complex logs or guesswork - we'll make it simple and intuitive for you to understand your resources.
- Insights and Analysis: We'll help you gain deep insights into your Kubernetes resources with powerful analytics and analysis. Get real-time information on resource utilization, health status, and configuration changes, so you can quickly identify and resolve issues, track changes over time, and optimize your resources for better performance.
- User-friendly Web Frontend: We believe in making KIT easy to use and visually appealing. Our web frontend is designed to be user-friendly and responsive, with clear representations of your inventory, including status indicators and configuration details. Find what you need, when you need it, with a clean and intuitive interface that fits seamlessly into your workflow.
- Integration with Kubernetes Ecosystem: We'll seamlessly integrate with the Kubernetes ecosystem, working smoothly with tools like kubectl, Kubernetes Dashboard, and Kubernetes-native logging solutions. You can continue using your existing authentication and authorization mechanisms, without any disruption to your workflow. We're here to enhance your existing Kubernetes setup, not complicate it.

With these goals in mind, our product mission is to simplify Kubernetes management and monitoring for DevOps teams, helping you optimize your resources and streamline your operations. We're committed to providing you with a user-friendly tool that empowers you to effectively manage your Kubernetes cluster and ensure the smooth operation of your containerized applications.

## Setup

To run the project please install [Docker and Docker compose](https://docs.docker.com/get-docker/) and execute `docker-compose up` in the root directory of the project.

If you also want to contribute please install the following:

- [golang](https://go.dev/doc/install)
- [NodeJS/npm](https://nodejs.org/en/download)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)

```bash
cd Explorer && npm install && cd .. && pre-commit install && pre-commit install -t commit-msg
```
