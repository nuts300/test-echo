# Test echo

## env
- mysql
- go 1.9.2

## Install

### Install dep
```
brew install dep
```

### Install dependencies
```
dep ensure
```

## Start app on development

### Install gin
```
go get github.com/codegangsta/gin
```

### Start db for dev
```
./start_db_dev.sh
```

### Start app for dev
```
./start_app_dev.sh
```

## Testing
```
./test.sh
```