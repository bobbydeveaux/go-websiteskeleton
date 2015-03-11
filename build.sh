cd static
r.js -o build.js
r.js -o cssIn=./css/style.css out=./dist/style.min.css optimizeCss=default
r.js -o cssIn=./css/style-large.css out=./dist/style-large.min.css optimizeCss=default
r.js -o cssIn=./css/style-xlarge.css out=./dist/style-xlarge.min.css optimizeCss=default
r.js -o cssIn=./css/style-small.css out=./dist/style-small.min.css optimizeCss=default
r.js -o cssIn=./css/style-xsmall.css out=./dist/style-xsmall.min.css optimizeCss=default
r.js -o cssIn=./css/style-medium.css out=./dist/style-medium.min.css optimizeCss=default
