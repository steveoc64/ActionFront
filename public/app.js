var Ratings = ['OldGuard','Guard','Grenadier','Elite','CrackLine','Veteran','Regular','Conscript','Landwehr','Militia','Rabble'];
var DrillBooks = ['Light Infantry','French','Prussian','Russian','Austrian','British','Old School','Conscript','Militia','Mob'];
var Equips = ['Musket','Carbine','Superior Musket','Poor Musket','Rifle','Bayonet Only','Pike'];
var SkirmishRatings = ['Superior','Excellent','Good','Average','Poor'];
var CavMoveTypes = ['Heavy','Medium','Light','Lancer'];
var GunneryClasses = [0,1,2,3];
var GunTypes = ['12pdr','9pdr','8pdr','6pdr','4pdr','3pdr','2pdr'];
var HWTypes = ['6"','5.5"','10pdr','18pdr L','9pdr L','7pdr'];

function socketUrl(s) {
    var l = window.location;
    return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname + (((l.port != 80) && (l.port != 443)) ? ":" + l.port : "") + l.pathname + s;
}

angular.module("app", ['ui.router', 'ngGrid'])
 .config(['$urlRouterProvider', '$stateProvider',function ($urlRouterProvider, $stateProvider) {
 	$urlRouterProvider.otherwise('/');
 	$stateProvider
 		.state('unittypes', {
 			url: '/unittypes',
 			templateUrl: 'unittypes.html',
 			controller: 'UnitTypesCtrl'
 		})
 		.state('unittypes.cavalerie', {
 			url: '/cavalerie',
 			templateUrl: 'unittypes.cavalerie.html',
 			controller: 'CavalryCtrl'
 		})
 		.state('unittypes.infanterie', {
 			url: '/infanterie',
 			templateUrl: 'unittypes.infanterie.html',
 			controller: 'InfantryCtrl'
 		})
 		.state('unittypes.artillerie', {
 			url: '/artillerie',
 			templateUrl: 'unittypes.artillerie.html',
 			controller: 'ArtilleryCtrl'
 		})
 		.state('unittypes.etat', {
 			url: '/etat',
 			templateUrl: 'unittypes.etat.html',
 			controller: 'EtatCtrl'
 		})
 		.state('unittypes.reglement', {
 			url: '/reglement',
 			templateUrl: 'drillbook.html',
 			controller: 'DrillBookCtrl'
 		});
 }])
.factory('DataSocket', ["$rootScope", function($rootScope) {
  var service = {};
  $rootScope.FilterValues = {};

  service.connect = function(scope) {
  	service.scope = scope;
  	service.Entity = scope.Entity;
  	service.Data = scope.Data;
   	service.initMessage = JSON.stringify({"Action":"List", "Entity":service.Entity});

    if(service.ws) { 
		service.ws.send(service.initMessage);
    	return; 
    }

    var ws = new WebSocket(socketUrl('GameData'));
    service.isOpen = false;
  
    ws.onmessage = function(e) {
    	console.log(e);
		var RxMsg = JSON.parse(e.data);
	
		if (RxMsg.Entity == service.Entity) {
			// console.log($scope)
			console.log("Msg ->", RxMsg);

			switch (RxMsg.Action) {
				case "List":
					service.Data = RxMsg.Data;	
					break;
				case "Update":
					var gotSome = false;
					var data = RxMsg.Data;

					// If the ID of the record exists, update the record in the dataset
					angular.forEach(service.Data, function(v,i){
						if (data["@id"] === v["@id"]) {
							console.log("Updating record at pos",i,"to",data);
							angular.copy(data,service.Data[i]);
							gotSome = true;
						}
					});

					// else if any of our records have a blank ID, overwrite that as the new record
					if (!gotSome) {
						console.log("Add New Record");
						angular.forEach(service.Data, function(v,i){
							if (v["@id"] === "0") {
								console.log("Overwriting Blank record at pos",i);
								angular.copy(data,service.Data[i]);
								gotSome = true;
							}
						});
					}
					// otherwise - just append to the list
					if (!gotSome) {
						console.log("Adding new record");
						service.Data.push(data);
					}			

			}
		}
	    //service.callback(e.data);
	    service.scope.Data = service.Data;
    	service.scope.updateFilters();
	    service.scope.$apply();
    }

	ws.onopen = function() {
		service.ws.send(service.initMessage)
		service.isOpen = true;
	}

	ws.onclose = function(e) {
		console.log("Socket closed ?", e);
		service.isOpen = false;
		service.ws = null;
	}

	ws.onerror = function(e) {
		console.log("Socket error ?", e);
	}

    service.ws = ws;
  }
 
  service.send = function(message) {
  	if (!service.isOpen) {
  		// if we are dead, try to reopen the connection on demand
  		service.connect(service.entity);
  	} else {
    	service.ws.send(message); 		
  	}
  }
 
  service.subscribe = function(callback) {
    service.callback = callback;
  }
 
  return service;
}])
.controller("UnitTypesCtrl", ["$scope", "$rootScope", function($scope, $rootScope){
	$scope.FilterUpdate = function() {
		$rootScope.FilterValues = {"Nation":$scope.Nation, "Year":$scope.Year, "Name":$scope.Name};
		$scope.$broadcast('FilterUpdate', $rootScope.FilterValues);
	}
}])
.controller("InfantryCtrl", ["$scope", "DataSocket", "$rootScope", function($scope, DataSocket, $rootScope){
	$scope.Data = [];
	$scope.FilteredData = [];
	$scope.Ratings = Ratings;
	$scope.DrillBooks = DrillBooks;
	$scope.Equips = Equips;
	$scope.SkirmishRatings = SkirmishRatings;
	$scope.title = "L'Infanterie";
	$scope.Entity = "Infantry";

	DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'FilteredData',
		enableCellSelection: true,
        enableRowSelection: false,
        enableCellEdit: true,
        showGroupPanel: false,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableRowReordering: false,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields: ['From'],
        	directions: ['asc']
        },
        filterOptions: $scope.filterOptions,
        groups: ['Nation'],

        columnDefs: [
           	{field:'Nation', width: 80}, 
           	{field:'From', width: 50}, 
           	{field:'To', width: 50}, 
        	{field:'Name', width: 160}, 
        	{field:'Rating', width: 100, editableCellTemplate: 'ratingTemplate.html'},
        	{field:'DrillBook', width: 100, editableCellTemplate: 'drillBookTemplate.html'},
        	{field:'Layout', width: 80},
        	{field:'Fire', width: 40},
        	{field:'Elite', width: 40},
        	{field:'Equip', width: 100, editableCellTemplate: 'equipTemplate.html'},
        	{field:'Skirmish', width: 100, editableCellTemplate: 'skirmishRatingTemplate.html'},
        	{field:'Street', width: 100, editableCellTemplate: 'streetRatingTemplate.html'},
        	{field:'Shock', width: 100, editableCellTemplate: 'shockTemplate.html'}
        ]
	};

	$scope.$on('FilterUpdate', function(e,data) {
		console.log('FilterUpdate event',data);
		$scope.updateFilters();
	});

	$scope.updateFilters = function() {
		$scope.FilteredData = [];
		angular.forEach($scope.Data, function(v,i){
			var ok = true;
			if ('Nation' in $rootScope.FilterValues) {
				if ($rootScope.FilterValues.Nation && v.Nation.indexOf($rootScope.FilterValues.Nation) == -1) {
					ok = false;
				}
			}
			if (ok && 'Name' in $rootScope.FilterValues) {
				if ($rootScope.FilterValues.Name && v.Name.indexOf($rootScope.FilterValues.Name) == -1) {
					ok = false;
				}
			}
			if (ok && 'Year' in $rootScope.FilterValues) {
				theYear = parseInt($rootScope.FilterValues.Year);
				yearFrom = parseInt(v.From);
				yearTo = parseInt(v.To);
				if (theYear < yearFrom || theYear > yearTo) {
					ok = false;
				}
			}
			if (ok) {
				$scope.FilteredData.push(v);
			}
		});
	}

	$scope.update = function(row) {
		console.log("InfUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.FilteredData.push({"@id": '0', Nation: '~ ??? ~'})
    }
	
}])
.controller("CavalryCtrl", ["$scope", "DataSocket", "$rootScope", function($scope, DataSocket, $rootScope){
	$scope.Data = [];
	$scope.FilteredData = [];
	$scope.Ratings = Ratings;
	$scope.CavMoveTypes = CavMoveTypes;
	$scope.SkirmishRatings = SkirmishRatings;
	$scope.title = "La Cavalerie";
	$scope.Entity = "Cavalry";

	DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'FilteredData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields: ['From'],
        	directions: ['asc']
        },
        groups: ['Nation'],

        columnDefs: [
           	{field:'Nation', width: 80}, 
           	{field:'From', width: 50}, 
           	{field:'To', width: 50}, 
        	{field:'Name', width: 160}, 
        	{field:'Rating', width: 100, editableCellTemplate: 'ratingTemplate.html'},
        	{field:'Shock', width: 80},
        	{field:'Squadrons', width: 80},
        	{field:'Move', width: 100, editableCellTemplate: 'cavMovesTemplate.html'},
        	{field:'Skirmish', width: 100, editableCellTemplate: 'skirmishRatingTemplate.html'}
        ]
	};

	$scope.update = function(row) {
		console.log("CavUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.$on('FilterUpdate', function(e,data) {
		console.log('FilterUpdate event',data);
		$scope.updateFilters();
	});

	$scope.updateFilters = function() {
		$scope.FilteredData = [];
		angular.forEach($scope.Data, function(v,i){
			var ok = true;
			if ('Nation' in $rootScope.FilterValues) {
				if ($rootScope.FilterValues.Nation && v.Nation.indexOf($rootScope.FilterValues.Nation) == -1) {
					ok = false;
				}
			}
			if (ok && 'Name' in $rootScope.FilterValues) {
				if ($rootScope.FilterValues.Name && v.Name.indexOf($rootScope.FilterValues.Name) == -1) {
					ok = false;
				}
			}
			if (ok && 'Year' in $rootScope.FilterValues) {
				theYear = parseInt($rootScope.FilterValues.Year);
				yearFrom = parseInt(v.From);
				yearTo = parseInt(v.To);
				if (theYear < yearFrom || theYear > yearTo) {
					ok = false;
				}
			}
			if (ok) {
				$scope.FilteredData.push(v);
			}
		});
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.FilteredData.push({"@id": '0', Nation: '~ ??? ~'})
    }
	
}])
.controller("ArtilleryCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.FilteredData = [];
	$scope.Ratings = Ratings;
	$scope.GunTypes = GunTypes;
	$scope.HWTypes = HWTypes;
	$scope.title = "L'Artillerie";
	$scope.Entity = "Artillery";

	DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'FilteredData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields: ['From'],
        	directions: ['asc']
        },
        groups: ['Nation'],

        columnDefs: [
           	{field:'Nation', width: 80}, 
           	{field:'From', width: 50}, 
           	{field:'To', width: 50}, 
        	{field:'Name', width: 160}, 
        	{field:'Rating', width: 100, editableCellTemplate: 'ratingTemplate.html'},
        	{field:'Class', width: 60},
        	{field:'Guns', width: 100, editableCellTemplate: 'gunTypeTemplate.html'},
        	{field:'HW', width: 100, editableCellTemplate: 'hwTemplate.html'},
        	{field:'Sections', width: 80},
        	{field:'Horse', width: 100}
        ]
	};

	$scope.$on('FilterUpdate', function(e,data) {
		console.log('FilterUpdate event',data);
		$scope.updateFilters();
	});

	$scope.updateFilters = function() {
		$scope.FilteredData = [];
		angular.forEach($scope.Data, function(v,i){
			var ok = true;
			if ('Nation' in $rootScope.FilterValues) {
				if ($rootScope.FilterValues.Nation && v.Nation.indexOf($rootScope.FilterValues.Nation) == -1) {
					ok = false;
				}
			}
			if (ok && 'Name' in $rootScope.FilterValues) {
				if ($rootScope.FilterValues.Name && v.Name.indexOf($rootScope.FilterValues.Name) == -1) {
					ok = false;
				}
			}
			if (ok && 'Year' in $rootScope.FilterValues) {
				theYear = parseInt($rootScope.FilterValues.Year);
				yearFrom = parseInt(v.From);
				yearTo = parseInt(v.To);
				if (theYear < yearFrom || theYear > yearTo) {
					ok = false;
				}
			}
			if (ok) {
				$scope.FilteredData.push(v);
			}
		});
	}

	$scope.update = function(row) {
		console.log("GunsUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.FilteredData.push({"@id": '0', Nation: '~ ??? ~'})
    }
	
}]);
