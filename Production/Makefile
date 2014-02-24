all: clean fetch production

clean:
	rm -f *.js *.css *.1 *.2
	echo All Clean
	ls -l

fetch: clean
	wget netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css 
	wget netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.min.css
	wget cdnjs.cloudflare.com/ajax/libs/ng-grid/2.0.7/ng-grid.min.css
	wget ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js
	wget netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js
	wget ajax.googleapis.com/ajax/libs/angularjs/1.2.10/angular.min.js
	wget cdnjs.cloudflare.com/ajax/libs/angular-ui-router/0.2.0/angular-ui-router.min.js
	wget cdnjs.cloudflare.com/ajax/libs/ng-grid/2.0.7/ng-grid.min.js
	echo All Fetched
	ls -l

production: 
	cat bootstrap.min.css font-awesome.min.css ng-grid.min.css ../public/app.css > ../public/gameDataEditor.css
	cat jquery.min.js bootstrap.min.js angular.min.js angular-ui-router.min.js ng-grid.min.js ../public/app.js | uglifyjs > ../public/gameDataEditor.js
	ls -l ../public/gameDataEditor.css ../public/gameDataEditor.js

