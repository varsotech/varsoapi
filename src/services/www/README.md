# VA Web

## Getting Started

1. Create a personal access token in GitHub (classic) - [click here](https://github.com/settings/tokens). Save the token in the Varso Bitwarden.
2. Login to GitHub's NPM server with the access token (use your GitHub username and email):

```
npm login --scope=@varsotech --auth-type=legacy --registry=https://npm.pkg.github.com
```

3. Add this to `~/.npmrc`:

```
@varsotech:registry=https://npm.pkg.github.com
```

3. Install packages

```
yarn install
```

4. Setup env

```
./scripts/setup_dev_env.sh
```

5. Start the react server

```
yarn dev
```

## Using VA API

When making changes to the VA server, the API might change and require you to use the new package.
To install it run the following command with the new package version:

```
npm install @varsotech/va@0.0.1
```
