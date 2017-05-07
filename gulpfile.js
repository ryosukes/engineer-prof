var path = require('path');
var gulp = require("gulp");
var sass = require("gulp-sass");
var plumber  = require('gulp-plumber');
var notify   = require('gulp-notify');
var cleanCSS = require('gulp-clean-css');
var rename   = require('gulp-rename');

var scssBasePath = 'assets/scss/';

gulp.task("scss", function() {
	return gulp.src(scssBasePath + '**/*.scss', {base: scssBasePath})
		.pipe(plumber({
			errorHandler: notify.onError("Error: <%= error.message %>")
		}))
		.pipe(sass())
		.pipe(cleanCSS())
		.pipe(rename({suffix: '.min'}))
		.pipe(gulp.dest('public/css/'));
});

gulp.task('watch', function() {
	gulp.watch([scssBasePath + '**/*.scss'], ['scss']);
});

gulp.task('default', ['scss']);
