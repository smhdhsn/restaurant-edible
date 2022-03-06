# Food

This repository is a simple food ordering system.

## About

Following functionalities are provided in this project.

- Store and manage food recipes.
- Submit orders.
- Update inventory's stocks after each order submittion.

### APIs

- `/api/menu` returns possible foods based on available raw materials.
- `/api/order` registers an order and updates inventory.

## Getting started

Make sure you have `docker` and `docker-compose` installed on your device.

### Starting the containers

Following command starts the application:

```shell
    make up
```

> :information_source: You need to wait a while for application to start.

Following command is used to access the container's command line:

```shell
    make bash
```

Following command creates sample food recipe inside database:

```shell
    make recipe
```

Following command stores missing food components inside database:

```shell
    make buy
```

Following command deletes useless food components from database:

```shell
    make recycle
```
