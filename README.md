<h1 align="center">
  <a href="https://afaqurk.github.io/linux-dash">
    <img src="https://raw.githubusercontent.com/afaqurk/screenshots/master/linux-dash/v2.0-logo.png"/>
  </a>
</h1>

<p align="center">
  <sub>v2.0</sub><br/>
  <small>A simple & low-overhead web dashboard for linux systems</small>
</p>

<p align="center">
  <small>
    <a href="https://afaqurk.github.io/linux-dash">Demo</a> &nbsp;|&nbsp;
    <a href="https://github.com/afaqurk/linux-dash/wiki">
      Docs
    </a>
  </small>
</p>

![image](https://user-images.githubusercontent.com/1940588/126266154-0349fb54-d83b-4980-9b54-3421a0ea4194.png)
![image](https://user-images.githubusercontent.com/1940588/126266213-f3a7db9a-a8ec-45f3-8ba3-df1012721196.png)
![image](https://user-images.githubusercontent.com/1940588/126266273-055fce1d-b91a-4c44-898b-784d1d921b7a.png)
![image](https://user-images.githubusercontent.com/1940588/126266302-ebe7fe13-403a-494e-b1d3-076a27aba646.png)
![image](https://user-images.githubusercontent.com/1940588/126266322-a8bd78c9-389e-4961-8971-6052c16abd91.png)

<p align="center">
  <a href="https://gitter.im/afaqurk/linux-dash">
    <img
      src="https://badges.gitter.im/gitterHQ/gitter.png"
      alt="linux-dash Gitter chat">
  </a>
</p>

<br/>

## Features
* **Small** ----- Under 400KB on disk _(with .git removed)!_
* **Simple** ---- A minimalist, beautiful dashboard
* **Easy** ------ Drop-in installation
* **Versatile** -- Choose your stack from Node.js, Go, Python, PHP

## Installation

### Step 1
```sh
## 1. clone the repo
git clone --depth 1 https://github.com/afaqurk/linux-dash.git

## 2. go to the cloned directory
cd linux-dash/app/server

```
OR, if you prefer to download manually:

```sh
## 1. Download the .zip
curl -LOk https://github.com/afaqurk/linux-dash/archive/master.zip && unzip master.zip

## 2. navigate to downloaded & unzipped dir
cd linux-dash-master/app/server

```

### Step 2

See instructions for preferred server linux-dash server _(all included)_:

* [Node.js](#if-using-nodejs) _(recommended)_
* [Go](#if-using-go)
* [Python](#if-using-python)
* [PHP](#if-using-php)

#### If Using Node.js
```sh
## install dependencies
npm install --production

## start linux-dash (on port 80 by default; may require sudo)
## You may change this with the `LINUX_DASH_SERVER_PORT` environment variable (eg. `LINUX_DASH_SERVER_PORT=8080 node server`)
## or provide a --port flag to the command below
## Additionally, the server will listen on every network interface (`0.0.0.0`).
## You may change this with the `LINUX_DASH_SERVER_HOST` environment variable (eg. `LINUX_DASH_SERVER_HOST=127.0.0.1 node server`)
## or provide a --host flag to the command below
node index.js

```

#### If Using Go
```sh
## start the server (on port 80 by default; may require sudo)
go run index.go
```

To build a binary, run `go build && ./server -h`. See [@tehbilly](https://github.com/sergeifilippov)'s notes [here](https://github.com/afaqurk/linux-dash/pull/281) for binary usage options

#### If Using Python
```sh
# Start the server (on port 80 by default; may require sudo).
python index.py
```

#### If Using PHP

1. Make sure you have the `exec`, `shell_exec`, and `escapeshellarg` functions enabled
2. Point your web server to `app/` directory under `linux-dash`
2. Restart your web server (Apache, nginx, etc.)
  - For PHP + Apache setup follow the [Digital Ocean tutorial](https://www.digitalocean.com/community/tutorials/how-to-install-linux-dash-on-ubuntu-14-04).
  - For help with nginx setup, see [this gist](https://gist.github.com/sergeifilippov/8909839) by [@sergeifilippov](https://github.com/sergeifilippov).

## Support

For general help, please use the [Gitter chat room](https://gitter.im/afaqurk/linux-dash).

## Security

**It is strongly recommended** that all linux-dash installations be protected via a security measure of your choice.

Linux Dash does not provide any security or authentication features.
