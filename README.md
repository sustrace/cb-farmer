# Github Contributions Farmer

## Step-by-step

**1. Generate two github tokens.**

- [Fine-grained access token](https://github.com/settings/tokens?type=beta)
- [Classic token](https://github.com/settings/tokens)

**2. Create and configure the `app.env` file in the config folder.** </br>
**Farmer will not work if you don`'t configure these fields:**

```
ACCESS_TOKEN=
CLASSIC_TOKEN=
USER_NAME=
USER_EMAIL=

START_DATE=
END_DATE=
```

Date format: "2022-01-01T00:00:00Z". <br />
Note: If ```START_DATE``` is not specified, the current date will be used as default.

It's not neccessarry to configure other parameters, but it's on your own.

```
REPOSITORIES_PATH=
FILE_NAME=
TARGET_REPOSITORY=
REPOSITORY_PREFIX=
```


## How-to-run

- Windows

```bash
make all-w
```

- Linux

```bash
make all-l
```
