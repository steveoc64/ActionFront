all: production

clean:
	rm -f *.min.js *.min.css *.1 *.2 *.js
	echo All Clean
	ls -l

fetch: clean
	wget netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css
	wget netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.min.css
	wget cdnjs.cloudflare.com/ajax/libs/ng-grid/2.0.7/ng-grid.min.css
	wget mgcrea.github.io/angular-strap/static/styles/angular-motion.min.css
	wget mgcrea.github.io/angular-strap/static/styles/bootstrap-additions.min.css
	wget mgcrea.github.io/angular-strap/static/styles/libraries.min.css

	wget ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js
	wget netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js
	wget code.angularjs.org/1.2.13/angular.min.js
	wget code.angularjs.org/1.2.13/angular-animate.min.js
	wget code.angularjs.org/1.2.13/angular-sanitize.min.js
	wget cdnjs.cloudflare.com/ajax/libs/angular-ui-router/0.2.0/angular-ui-router.min.js
	wget mgcrea.github.io/angular-strap/dist/angular-strap.js
	wget mgcrea.github.io/angular-strap/dist/angular-strap.tpl.js
	echo All Fetched
	ls -l

production: 
	cat bootstrap.min.css font-awesome.min.css ng-grid.min.css angular-motion.min.css bootstrap-additions.css libraries.min.css ../public/app.css > ../public/gameDataEditor.css
	cat bootstrap.css font-awesome.min.css ng-grid.min.css angular-motion.min.css bootstrap-additions.css libraries.min.css ../public/app.css > ../public/gameDataEditor.css

	cat jquery.min.js bootstrap.min.js angular.min.js angular-animate.min.js angular-sanitize.min.js angular-ui-router.min.js angular-strap.js angular-strap.tpl.js ng-grid.js.local ../public/app.js | uglifyjs > ../public/gameDataEditor.js
	ls -l ../public/gameDataEditor.css ../public/gameDataEditor.js
