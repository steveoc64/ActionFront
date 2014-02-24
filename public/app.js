var Ratings = ['OldGuard','Guard','Grenadier','Elite','CrackLine','Veteran','Regular','Conscript','Landwehr','Militia','Rabble'];
var DrillBooks = ['Light Infantry','French','Prussian','Russian','Austrian','British','Old School','Conscript','Militia','Mob'];
var Equips = ['Musket','Carbine','Superior Musket','Poor Musket','Rifle','Bayonet Only','Pike'];
var SkirmishRatings = ['Superior','Excellent','Good','Average','Poor'];
var CavMoveTypes = ['Heavy','Medium','Light','Lancer'];
var GunneryClasses = [0,1,2,3];
var GunTypes = ['12pdr','9pdr','8pdr','6pdr','4pdr','3pdr','2pdr'];
var HWTypes = ['6"','5.5"','10pdr','18pdr L','9pdr L','7pdr'];
var MEOrders = ['Attack','Defend','Bombard','Support/Intercept','March','Rest','Redeploy','BreakOff','Screen','RearGuard'];
var StaffRatings = ['Good','Average','Poor'];

var defaultGridOptions = { 
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
        groupsCollapsedByDefault: false
};

function socketUrl(s) {
    var l = window.location;
    return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname + (((l.port != 80) && (l.port != 443)) ? ":" + l.port : "") + s;
}

function contains(haystack,needle) {
	if (typeof needle != 'string') {
		return true;
	}
	if (typeof haystack != 'string') {
		return true;
	}

	var h1 = haystack.toLowerCase();
	var n1 = needle.toLowerCase();
	return (h1.indexOf(n1) != -1);
}

angular.module("app", ['ui.router', 'ngGrid'])
 .config(['$urlRouterProvider', '$stateProvider',function ($urlRouterProvider, $stateProvider) {
 	$urlRouterProvider.otherwise('/');
 	$stateProvider
 		.state('unitTypes', {
 			url: '/unitTypes',
 			templateUrl: 'unitTypes.html',
 			controller: 'UnitTypesCtrl'
 		})
 		.state('unitTypes.formations', {
 			url: '/formations',
 			templateUrl: 'unitTypes.formations.html',
 			controller: 'FormationsCtrl'
 		})
 		.state('unitTypes.cavalerie', {
 			url: '/cavalerie',
 			templateUrl: 'unitTypes.cavalerie.html',
 			controller: 'CavalryCtrl'
 		})
 		.state('unitTypes.infanterie', {
 			url: '/infanterie',
 			templateUrl: 'unitTypes.infanterie.html',
 			controller: 'InfantryCtrl'
 		})
 		.state('unitTypes.artillerie', {
 			url: '/artillerie',
 			templateUrl: 'unitTypes.artillerie.html',
 			controller: 'ArtilleryCtrl'
 		})
 		.state('unitTypes.etat', {
 			url: '/etat',
 			templateUrl: 'unitTypes.etat.html',
 			controller: 'EtatCtrl'
 		})
 		.state('unitTypes.reglement', {
 			url: '/reglement',
 			templateUrl: 'unitTypes.reglement.html',
 			controller: 'DrillBookCtrl'
 		})
 		.state('unitTypes.equip', {
 			url: '/equip',
 			templateUrl: 'unitTypes.equip.html',
 			controller: 'EquipCtrl'
 		})
 		.state('initTables', {
 			url: '/initTables',
 			templateUrl: 'initTables.html',
 			controller: 'InitTablesCtrl'
 		})
 		.state('corpsOrders', {
 			url: '/corpsOrders',
 			templateUrl: 'corpsOrders.html',
 			controller: 'CorpsOrdersCtrl'
 		})
 		.state('meOrders', {
 			url: '/meOrders',
 			templateUrl: 'meOrders.html',
 			controller: 'MEOrdersCtrl'
 		})
		.state('orderArrival', {
 			url: '/orderArrival',
 			templateUrl: 'orderArrival.html',
 			controller: 'OrderArrivalCtrl'
 		})
 		.state('orderActivation', {
 			url: '/orderActivation',
 			templateUrl: 'orderActivation.html',
 			controller: 'OrderActivationCtrl'
 		})
 		.state('commanderAction', {
 			url: '/commanderAction',
 			templateUrl: 'commanderAction.html',
 			controller: 'CommanderActionCtrl'
 		})
 		.state('MEMorale', {
 			url: '/MEMorale',
 			templateUrl: 'MEMorale.html',
 			controller: 'MEMoraleCtrl'
 		})
 		.state('MEPanic', {
 			url: '/MEPanic',
 			templateUrl: 'MEPanic.html',
 			controller: 'MEPanicCtrl'
 		})
 		.state('UnitMorale', {
 			url: '/UnitMorale',
 			templateUrl: 'UnitMorale.html',
 			controller: 'UnitMoraleCtrl'
 		})
 		.state('FireDisc', {
 			url: '/FireDics',
 			templateUrl: 'FireDisc.html',
 			controller: 'FireDiscCtrl'
 		})
 		.state('InitBadMorale', {
 			url: '/InitBadMorale',
 			templateUrl: 'InitBadMorale.html',
 			controller: 'InitBadMoraleCtrl'
 		})
 		.state('BonusImpulse', {
 			url: '/BonusImpulse',
 			templateUrl: 'BonusImpulse.html',
 			controller: 'BonusImpulseCtrl'
 		})
 		.state('MEFatigue', {
 			url: '/MEFatigue',
 			templateUrl: 'MEFatigue.html',
 			controller: 'MEFatigueCtrl'
 		})
 		.state('FatigueRecovery', {
 			url: '/FatigueRecovery',
 			templateUrl: 'FatigueRecovery.html',
 			controller: 'FatigueRecoveryCtrl'
 		})
 		.state('MoraleRecovery', {
 			url: '/MoraleRecovery',
 			templateUrl: 'MoraleRecovery.html',
 			controller: 'MoraleRecoveryCtrl'
 		})
 		;
 }])
.factory('DataSocket', ["$rootScope", function($rootScope) {
  var service = {};
  $rootScope.FilterValues = {};

  service.connect = function(subscriptions) {
  	//service.scope = scope;
  	//service.Entity = scope.Entity;
  	//service.Data = scope.Data;

  	// subscriptions is an array of objects in the form
  	// Entity, Data, Callback
  	service.subscriptions = subscriptions;

  	//console.log("Subs = ",subscriptions);
   	//service.initMessage = JSON.stringify({"Action":"List", "Entity":service.Entity});

    if(service.ws) { 
    	angular.forEach(service.subscriptions, function(v,i){
    		initMsg = JSON.stringify({"Action":"List", "Entity":v.Entity});
    		service.ws.send(initMsg);
    	});
		//service.ws.send(service.initMessage);
    	return; 
    }

    var ws = new WebSocket(socketUrl('/GameData'));
    service.isOpen = false;
  
    ws.onmessage = function(e) {
    	//console.log(e);
		var RxMsg = JSON.parse(e.data);
	
		angular.forEach(service.subscriptions, function(sub,isub) {
			if (RxMsg.Entity == sub.Entity) {
				// console.log($scope)
				console.log("Msg ->", RxMsg);
				var gotSome = false;

				switch (RxMsg.Action) {
					case "List":
						//sub.Data = RxMsg.Data;	
						angular.copy(RxMsg.Data,sub.Data);
						gotSome = true;
						break;
					case "Update":
						var data = RxMsg.Data;

						// If the ID of the record exists, update the record in the dataset
						angular.forEach(sub.Data, function(v,i){
							if (data["@id"] === v["@id"]) {
								console.log("Updating record at pos",i,"to",data);
								angular.copy(data,sub.Data[i]);
								gotSome = true;
							}
						});

						// else if any of our records have a blank ID, overwrite that as the new record
						if (!gotSome) {
							console.log("Add New Record", data);
							angular.forEach(sub.Data, function(v,i){
								if (v["@id"] === "0") {
									console.log("Overwriting Blank record at pos",i);
									angular.copy(data,sub.Data[i]);
									gotSome = true;
								}
							});
						}
						// otherwise - just append to the list
						if (!gotSome) {
							console.log("Adding new record",data);
							sub.Data.push(data);
						}			
				} // switch
				if (gotSome) {
					sub.Callback(sub.Data);
				}
			} // if msg entity = this entity
		});
    }

	ws.onopen = function() {
    	angular.forEach(service.subscriptions, function(v,i){
    		initMsg = JSON.stringify({"Action":"List", "Entity":v.Entity});
    		service.ws.send(initMsg);
    	});
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
  		service.connect(service.subscriptions);
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
.controller("FormationsCtrl", ["$scope", "DataSocket", "$rootScope", function($scope, DataSocket, $rootScope){
	$scope.Data = [];
	$scope.FilteredData = [];
	$scope.title = "National Organisations";
	$scope.docs = "Table 1.2";
	$scope.Entity = "NationalOrg";

	//DataSocket.connect($scope);

	$scope.gridOptions = angular.copy(defaultGridOptions);
    $scope.gridOptions.sortInfo = {
        	fields: ['From'],
        	directions: ['asc']
        };
    $scope.gridOptions.groups = ['Nation'];
    $scope.gridOptions.columnDefs = [
           	{field:'Nation', width: 120, visible: false}, 
           	{field:'From', width: 60}, 
           	{field:'To', width: 60}, 
        	{field:'InfantryME', width: 300}, 
        	{field:'CavalryME', width: 300},
        	{field:'Corps', width: 300}
        ];

	$scope.$on('FilterUpdate', function(e,data) {
		console.log('FilterUpdate event',data);
		$scope.updateFilters();
	});

	$scope.updateFilters = function() {
		$scope.FilteredData = [];
		angular.forEach($scope.Data, function(v,i){
			var ok = true;
			if ('Nation' in $rootScope.FilterValues) {
				if (!contains(v.Nation,$rootScope.FilterValues.Nation)) {
					ok = false;
				}
			}
			if (ok && 'Name' in $rootScope.FilterValues) {
				var needle = $rootScope.FilterValues.Name;
				if (!contains(v.InfantryME,needle) &&
					!contains(v.CavalryME,needle) &&
					!contains(v.Corps,needle)) {
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
		console.log("FormationUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.FilteredData.push({"@id": '0', Nation: '~ ??? ~', From: 1792, To: 1815})
    }

    $scope.changeData = function(d) {
    	$scope.updateFilters();
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);		
}])
.controller("InfantryCtrl", ["$scope", "DataSocket", "$rootScope", function($scope, DataSocket, $rootScope){
	$scope.Data = [];
	$scope.FilteredData = [];
	$scope.Ratings = Ratings;
	$scope.DrillBooks = DrillBooks;
	$scope.Equips = Equips;
	$scope.SkirmishRatings = SkirmishRatings;
	$scope.title = "L'Infanterie";
	$scope.docs = "Appendix C";
	$scope.Entity = "Infantry";

	//DataSocket.connect($scope);

	$scope.gridOptions = angular.copy(defaultGridOptions);
    $scope.gridOptions.sortInfo = {
        	fields: ['From'],
        	directions: ['asc']
        };
    $scope.gridOptions.groups = ['Nation'];
    $scope.gridOptions.columnDefs = [
           	{field:'Nation', width: 80, visible: false}, 
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
        ];

	$scope.$on('FilterUpdate', function(e,data) {
		console.log('FilterUpdate event',data);
		$scope.updateFilters();
	});

	$scope.updateFilters = function() {
		$scope.FilteredData = [];
		angular.forEach($scope.Data, function(v,i){
			var ok = true;
			if ('Nation' in $rootScope.FilterValues) {
				if (!contains(v.Nation,$rootScope.FilterValues.Nation)) {
					ok = false;
				}
			}
			if (ok && 'Name' in $rootScope.FilterValues) {
				var needle = $rootScope.FilterValues.Name;
				if (!contains(v.Name,needle) &&
					!contains(v.Rating,needle)) {
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

    $scope.changeData = function(d) {
    	$scope.updateFilters();
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);	
	
}])
.controller("CavalryCtrl", ["$scope", "DataSocket", "$rootScope", function($scope, DataSocket, $rootScope){
	$scope.Data = [];
	$scope.FilteredData = [];
	$scope.Ratings = Ratings;
	$scope.CavMoveTypes = CavMoveTypes;
	$scope.SkirmishRatings = SkirmishRatings;
	$scope.title = "La Cavalerie";
	$scope.docs = "Appendix B";
	$scope.Entity = "Cavalry";

	//DataSocket.connect($scope);

	$scope.gridOptions = angular.copy(defaultGridOptions);
    $scope.gridOptions.sortInfo = {
        	fields: ['From'],
        	directions: ['asc']
        };
    $scope.gridOptions.groups = ['Nation'];
    $scope.gridOptions.columnDefs = [
           	{field:'Nation', width: 80, visible: false}, 
           	{field:'From', width: 50}, 
           	{field:'To', width: 50}, 
        	{field:'Name', width: 160}, 
        	{field:'Rating', width: 100, editableCellTemplate: 'ratingTemplate.html'},
        	{field:'Shock', width: 80},
        	{field:'Squadrons', width: 80},
        	{field:'Move', width: 100, editableCellTemplate: 'cavMovesTemplate.html'},
        	{field:'Skirmish', width: 100, editableCellTemplate: 'skirmishRatingTemplate.html'}
        ];

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
				if (!contains(v.Nation,$rootScope.FilterValues.Nation)) {
					ok = false;
				}
			}
			if (ok && 'Name' in $rootScope.FilterValues) {
				if (!contains(v.Name,$rootScope.FilterValues.Name)) {
					ok = false;
				}
			}
			if (ok && 'Name' in $rootScope.FilterValues) {
				var needle = $rootScope.FilterValues.Name;
				if (!contains(v.Name,needle) &&
					!contains(v.Rating,needle)) {
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

    $scope.changeData = function(d) {
    	$scope.updateFilters();
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);	
	
}])
.controller("ArtilleryCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.FilteredData = [];
	$scope.Ratings = Ratings;
	$scope.GunTypes = GunTypes;
	$scope.HWTypes = HWTypes;
	$scope.title = "L'Artillerie";
	$scope.docs = "Appendix D";
	$scope.Entity = "Artillery";

	//DataSocket.connect($scope);

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
        groupsCollapsedByDefault: false,

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
        	{field:'Horse', width: 100, editableCellTemplate: 'horseArtyTemplate.html'}

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
				if (!contains(v.Nation,$rootScope.FilterValues.Nation)) {
					ok = false;
				}
			}
			if (ok && 'Name' in $rootScope.FilterValues) {
				var needle = $rootScope.FilterValues.Name;
				if (!contains(v.Name,needle) &&
					!contains(v.Guns,needle) &&
					!contains(v.HW,needle)) {
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

    $scope.changeData = function(d) {
    	$scope.updateFilters();
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);	
	
}])
.controller("EtatCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.FilteredData = [];
	$scope.StaffRatings = StaffRatings;
	$scope.title = "L'Etat Major";
	$scope.docs = "Appendix G";
	$scope.Entity = "EtatMajor";

	//DataSocket.connect($scope);

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
        groupsCollapsedByDefault: false,

        columnDefs: [
           	{field:'Nation', width: 250}, 
           	{field:'From', width: 50}, 
           	{field:'To', width: 50}, 
        	{field:'Rating', width: 120, editableCellTemplate: 'staffRatingTemplate.html'},
        	{field:'Value', width: 60}
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
				if (!contains(v.Nation,$rootScope.FilterValues.Nation)) {
					ok = false;
				}
			}
			if (ok && 'Name' in $rootScope.FilterValues) {
				var needle = $rootScope.FilterValues.Name;
				if (!contains(v.Rating,needle)) {
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
		console.log("StaffUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.FilteredData.push({"@id": '0', Nation: '~ ??? ~'})
    }

    $scope.changeData = function(d) {
    	$scope.updateFilters();
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);		
}])
.controller("DrillBookCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.title = "L'Reglement / Drill Book";
	$scope.docs = "";
	$scope.Entity = "Drill";

	//DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: false,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields: ['Name'],
        	directions: ['asc']
        },
        columnDefs: [
           	{field:'Name', width: 200}, 
            {field:'Entries', 
           	 width: 500, 
           	 displayName: 'Formations',
           	 editableCellTemplate: 'meOrdersTemplate.html', 
           	 cellTemplate: 'reglementTable.html'
           	},

        ]
	};

	$scope.update = function(row) {
		console.log("DrillUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.FilteredData.push({"@id": '0', Name: '~ ??? ~'})
    }

    $scope.changeData = function(d) {
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);		
}])
.controller("EquipCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.title = "Equipment Types";
	$scope.docs = "";
	$scope.Entity = "Equip";

	//DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'Data',
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
        	fields: ['Name'],
        	directions: ['asc']
        },
        columnDefs: [
           	{field:'Name', width: 200}, 
           	{field:'SK', displayName: 'Skirmish', width: 200}, 
           	{field:'Volley', width: 200}, 
           	{field:'Close', width: 200}, 
           	{field:'Long', width: 200}, 
        ]
	};

	$scope.update = function(row) {
		console.log("EquipUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.FilteredData.push({"@id": '0', Name: '~ ??? ~'})
    }

    $scope.changeData = function(d) {
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);		
}])
.controller("InitTablesCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.title = "Initiative Modifiers";
	$scope.docs = "Table 11.1";
	$scope.Entity = "InitTable";

	//DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'Data',
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
        	fields: ['Value'],
        	directions: ['desc']
        },
        columnDefs: [
           	{field:'Factor', width: 300}, 
           	{field:'Value', width: 50}
        ]
	};

	$scope.update = function(row) {
		console.log("InitUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.updateFilters = function() {
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Factor: '~ ??? ~'})
    }

    $scope.changeData = function(d) {
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);	
	
}])
.controller("CorpsOrdersCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.MEOrders = MEOrders;
	$scope.title = "Corps Orders";
	$scope.docs = "Table 3.1";
	$scope.Entity = "CorpsOrder";

	//DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        columnDefs: [
           	{field:'Order', width: 100}, 
           	{field:'MEOrders', 
           	 width: 500, 
           	 displayName: 'Allowed ME Orders',
           	 editableCellTemplate: 'meOrdersTemplate.html', 
           	 cellTemplate: '<div class="ngBigCell"><ul><li ng-repeat="i in row.entity.MEOrders">{{i}}</li></ul></div>'
           	},
           	{field: 'Stipulation', width: 400}
        ]
	};

	$scope.update = function(row) {
		console.log("CorpsOrderUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.updateFilters = function() {
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Order: '~ ??? ~', MEOrders: ['Attack','Defend','Withdraw'], Stipulation: 'Enter constraints for the Corps under this order'})
    }

    $scope.changeData = function(d) {
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);	

}])
.controller("MEOrdersCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.MEOrders = MEOrders;
	$scope.title = "ME Orders";
	$scope.docs = "Table 4.1";
	$scope.Entity = "MEOrder";

	//DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        columnDefs: [
           	{field:'Order', width: 120}, 
           	{field:'Purpose', width: 400},
           	{field:'Notes', width: '40%'},
           	{field:'IfNonEngaged', width: 100, editableCellTemplate: 'ifNonEngagedTemplate.html'},
           	{field:'IfEngaged', width: 100, editableCellTemplate: 'ifEngagedTemplate.html'},
           	{field:'CavalryOnly', width: 100, editableCellTemplate: 'cavalryOnlyTemplate.html'},
           	{field:'DefendIfEngaged', width: 100, editableCellTemplate: 'defendIfEngagedTemplate.html'},
           	{field:'ShakenIfEngaged', width: 100, editableCellTemplate: 'shakenIfEngagedTemplate.html'}
        ]
	};

	$scope.update = function(row) {
		console.log("MEOrderUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.updateFilters = function() {
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Order: '~ ??? ~'})
    }

    $scope.changeData = function(d) {
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);

}])
.controller("OrderArrivalCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.title = "Order Arrival Calculation";
	$scope.docs = "Table 3.3";
	$scope.Entity = "OrderArrival";

	//DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'Data',
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
           	fields: ['Grids'],
        	directions: ['asc']    	
        },
        columnDefs: [
           	{field:'Grids', width: 120}, 
           	{field:'Delay', width: 400},
           	{field:'DGrids', displayName:'Diagonal Grids', width: 300}
        ]
	};

	$scope.update = function(row) {
		console.log("CorpsOrderUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.updateFilters = function() {
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Grids: '~ ??? ~'})
    }

    $scope.changeData = function(d) {
    	$scope.$apply();
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);
	
}])
.controller("OrderActivationCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.maintitle = "Order Activation";
	$scope.modtitle = "Activation Modifiers";
	$scope.docs = "Table 8.1";
	$scope.moddocs = "Table 8.1B";
	$scope.Entity = "OrderActivation";
	$scope.ModEntity = "OrderActivationMod";

	$scope.gridOptions = { 
		data: 'Data',
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
        	fields:['Dice'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Dice', width: 120}, 
           	{field:'Points', displayName: 'Activation Points', width: 200}
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 120}, 
           	{field:'Descr', displayName:'Description',width: 300},
           	{field:'Points', displayName: 'ME', width: 60},
           	{field:'CorpsPoints', displayName: 'Corps', width: 80}
        ]
	};

	$scope.update = function(row) {
		console.log("CorpsOrderUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}


	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		console.log(row);
		targetID = row["@id"];
		console.log("Looking for ID ",targetID);
		gotSome = false;
		angular.forEach($scope.Data, function(v,i){
			if (v["@id"] == targetID) {
				//console.log("The update is on the first grid");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
				gotSome = true;
			}
		});
		if (!gotSome) { 
				angular.forEach($scope.ModData, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the mod data grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
					gotSome = true;
				}
			})
		}
		if (!gotSome) {
			if ('Code' in row) {
				//console.log("The update is on the mod data grid because it has a property called Code");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
			} else {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Dice: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.ModData.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		//console.log("Change Data Callback",d);
		$scope.$apply();
	}

	$scope.changeModData = function(d) {
		//console.log("Change Mod Data Callback",d)
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.ModEntity, Data: $scope.ModData, Callback: $scope.changeModData}
	]);
	
}])
.controller("CommanderActionCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.title = "Commander Actions";
	$scope.docs = "Table 7.1, Table 12.1, Table 12.2";
	$scope.title2 = "Commander Action Score";
	$scope.docs2 = "Table 12.3 , Table 12.3A, Commander Ratings Apply.  +/-3 if leader is attached to a unit."
	$scope.Entity = "CommanderAction";
	$scope.Entity2 = "CAScore";

	//DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'Data',
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
           	fields: ['Cost'],
        	directions: ['asc']    	
        },
        groups: ['Who'],
        groupsCollapsedByDefault: false,
        columnDefs: [
        	{field:'Who',visible: false},
           	{field:'Code', width: 80}, 
           	{field:'Action', width: 350},
           	{field:'Cost', width: 60}
        ]
	};
	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
           	fields: ['Code'],
        	directions: ['asc']    	
        },
        columnDefs: [
        	{field:'Code',width:60},
           	{field:'Descr', displayName:'Description',width: 200}, 
           	{field:'A1', displayName:'1st',width: 80},
           	{field:'A2', displayName:'2nd',width: 80},
           	{field:'A3', displayName:'3rd',width: 80},
           	{field:'A4', displayName:'4th',width: 80}
        ]
	};

	$scope.update = function(row) {
		console.log("CAUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.updateFilters = function() {
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		console.log(row);
		targetID = row["@id"];
		console.log("Looking for ID ",targetID);
		gotSome = false;
		angular.forEach($scope.Data, function(v,i){
			if (v["@id"] == targetID) {
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
				gotSome = true;
			}
		});
		if (!gotSome) { 
				angular.forEach($scope.Data2, function(v,i){
				if (v["@id"] == targetID) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
					gotSome = true;
				}
			})
		}
		if (!gotSome) {
			if ('Descr' in row) {
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
			} else {
				$scope.update(evt.targetScope.row);	
			}
		}

    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Who: 'Corps',Code: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', Code: '~ ??? ~', Descr: 'Add new one here'})
    }

    $scope.changeData = function(d) {
    	$scope.$apply();
    }
    $scope.changeData2 = function(d) {
    	$scope.$apply();
    }

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.ModData, Callback: $scope.changeData2}
	]);
	
}])
.controller("MEMoraleCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.maintitle = "ME Morale Test";
	$scope.modtitle = "ME Morale Modifiers";
	$scope.docs = "Table 5.1";
	$scope.moddocs = "Table 5.1A";
	$scope.Entity = "MEMoraleTest";
	$scope.ModEntity = "MEMoraleMod";

	$scope.gridOptions = { 
		data: 'Data',
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
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Score', width: 80}, 
           	{field:'Descr', displayName: 'Description', width: 600}, 
           	{field:'Broken', width: 80, editableCellTemplate: 'moraleBrokenTemplate.html'}, 
           	{field:'Retreat', width: 80, editableCellTemplate: 'moraleRetreatTemplate.html'}, 
           	{field:'Shaken', width: 80, editableCellTemplate: 'moraleShakenTemplate.html'}, 
           	{field:'Steady', width: 80, editableCellTemplate: 'moraleSteadyTemplate.html'}, 
           	{field:'Fatigue', width: 80}, 
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 120}, 
           	{field:'Descr', displayName:'Description',width: 300},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("MEMoraleUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}


	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		console.log(row);
		targetID = row["@id"];
		console.log("Looking for ID ",targetID);
		gotSome = false;
		angular.forEach($scope.Data, function(v,i){
			if (v["@id"] == targetID) {
				//console.log("The update is on the first grid");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
				gotSome = true;
			}
		});
		if (!gotSome) { 
				angular.forEach($scope.ModData, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the mod data grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
					gotSome = true;
				}
			})
		}
		if (!gotSome) {
			if ('Code' in row) {
				//console.log("The update is on the mod data grid because it has a property called Code");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
			} else {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Score: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.ModData.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		//console.log("Change Data Callback",d);
		$scope.$apply();
	}

	$scope.changeModData = function(d) {
		//console.log("Change Mod Data Callback",d)
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.ModEntity, Data: $scope.ModData, Callback: $scope.changeModData}
	]);
	
}])
.controller("MEPanicCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.maintitle = "ME Panic Test";
	$scope.modtitle = "ME Panic Modifiers";
	$scope.docs = "Table 6.1";
	$scope.moddocs = "Table 6.1A";
	$scope.Entity = "MEPanicTest";
	$scope.ModEntity = "MEPanicMod";

	$scope.gridOptions = { 
		data: 'Data',
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
        	fields:['Broken'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Rating', width: 160}, 
           	{field:'Broken', width: 80}, 
           	{field:'Shaken', width: 80}, 
           	{field:'CarryOn', width: 80}, 
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 120}, 
           	{field:'Descr', displayName:'Description',width: 300},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("MEPanicUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}


	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		console.log(row);
		targetID = row["@id"];
		console.log("Looking for ID ",targetID);
		gotSome = false;
		angular.forEach($scope.Data, function(v,i){
			if (v["@id"] == targetID) {
				//console.log("The update is on the first grid");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
				gotSome = true;
			}
		});
		if (!gotSome) { 
				angular.forEach($scope.ModData, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the mod data grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
					gotSome = true;
				}
			})
		}
		if (!gotSome) {
			if ('Code' in row) {
				//console.log("The update is on the mod data grid because it has a property called Code");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
			} else {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Score: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.ModData.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		//console.log("Change Data Callback",d);
		$scope.$apply();
	}

	$scope.changeModData = function(d) {
		//console.log("Change Mod Data Callback",d)
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.ModEntity, Data: $scope.ModData, Callback: $scope.changeModData}
	]);
	
}])
.controller("UnitMoraleCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.maintitle = "Unit Morale Test";
	$scope.modtitle = "Unit Morale Modifiers";
	$scope.docs = "Table 19.1";
	$scope.moddocs = "Tables 19.1A - 19.1B";
	$scope.Entity = "UnitMoraleTest";
	$scope.ModEntity = "UnitMoraleMod";

	$scope.gridOptions = { 
		data: 'Data',
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
        	fields:['Pass'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Rating', width: 160}, 
           	{field:'Pass', width: 80}, 
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 120}, 
           	{field:'Descr', displayName:'Description',width: 300},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("UnitMoraleUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}


	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		console.log(row);
		targetID = row["@id"];
		console.log("Looking for ID ",targetID);
		gotSome = false;
		angular.forEach($scope.Data, function(v,i){
			if (v["@id"] == targetID) {
				//console.log("The update is on the first grid");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
				gotSome = true;
			}
		});
		if (!gotSome) { 
				angular.forEach($scope.ModData, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the mod data grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
					gotSome = true;
				}
			})
		}
		if (!gotSome) {
			if ('Code' in row) {
				//console.log("The update is on the mod data grid because it has a property called Code");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
			} else {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Score: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.ModData.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		//console.log("Change Data Callback",d);
		$scope.$apply();
	}

	$scope.changeModData = function(d) {
		//console.log("Change Mod Data Callback",d)
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.ModEntity, Data: $scope.ModData, Callback: $scope.changeModData}
	]);
	
}])
.controller("FireDiscCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.maintitle = "Fire Discipline Test";
	$scope.modtitle = "Fire Discipline Modifiers";
	$scope.docs = "Table 13.3";
	$scope.moddocs = "Tables 13.3A";
	$scope.Entity = "FireDisciplineTest";
	$scope.ModEntity = "FireDisciplineMod";

	$scope.gridOptions = { 
		data: 'Data',
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
        	fields:['Pass'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Rating', width: 160}, 
           	{field:'Pass', width: 80}, 
           	{field:'Fire', width: 80}, 
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 120}, 
           	{field:'Descr', displayName:'Description',width: 300},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("FireDiscUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}


	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		console.log(row);
		targetID = row["@id"];
		console.log("Looking for ID ",targetID);
		gotSome = false;
		angular.forEach($scope.Data, function(v,i){
			if (v["@id"] == targetID) {
				//console.log("The update is on the first grid");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
				gotSome = true;
			}
		});
		if (!gotSome) { 
				angular.forEach($scope.ModData, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the mod data grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
					gotSome = true;
				}
			})
		}
		if (!gotSome) {
			if ('Code' in row) {
				//console.log("The update is on the mod data grid because it has a property called Code");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
			} else {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Score: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.ModData.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		//console.log("Change Data Callback",d);
		$scope.$apply();
	}

	$scope.changeModData = function(d) {
		//console.log("Change Mod Data Callback",d)
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.ModEntity, Data: $scope.ModData, Callback: $scope.changeModData}
	]);
	
}])
.controller("InitBadMoraleCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.maintitle = "Initial Bad Morale Test";
	$scope.modtitle = "Initial Bad Morale Modifiers";
	$scope.docs = "Table 22.2";
	$scope.moddocs = "Tables 22.2A";
	$scope.Entity = "InitialBadMorale";
	$scope.ModEntity = "InitialBadMod";

	$scope.gridOptions = { 
		data: 'Data',
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
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Score', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 300}, 
           	{field:'Hits', width: 80}, 
           	{field:'Stay', width: 80, editableCellTemplate: 'stayTemplate.html'}
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 120}, 
           	{field:'Descr', displayName:'Description',width: 300},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("InitBadUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}


	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		console.log(row);
		targetID = row["@id"];
		console.log("Looking for ID ",targetID);
		gotSome = false;
		angular.forEach($scope.Data, function(v,i){
			if (v["@id"] == targetID) {
				//console.log("The update is on the first grid");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
				gotSome = true;
			}
		});
		if (!gotSome) { 
				angular.forEach($scope.ModData, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the mod data grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
					gotSome = true;
				}
			})
		}
		if (!gotSome) {
			if ('Code' in row) {
				//console.log("The update is on the mod data grid because it has a property called Code");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
			} else {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Score: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.ModData.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		//console.log("Change Data Callback",d);
		$scope.$apply();
	}

	$scope.changeModData = function(d) {
		//console.log("Change Mod Data Callback",d)
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.ModEntity, Data: $scope.ModData, Callback: $scope.changeModData}
	]);
	
}])
.controller("BonusImpulseCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.maintitle = "Bonus Impulse Test";
	$scope.modtitle = "Bonus Impulse Modifiers";
	$scope.docs = "Table 20.1";
	$scope.moddocs = "Tables 20.1A";
	$scope.Entity = "BonusImpulse";
	$scope.ModEntity = "BonusImpulseMod";

	$scope.gridOptions = { 
		data: 'Data',
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
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Score', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 300}, 
           	{field:'Another', displayName:'Bonus',width: 80, editableCellTemplate: 'anotherTemplate.html'},
           	{field:'Fatigue', width: 80, editableCellTemplate: 'fatigueTemplate.html'},
           	{field:'OneUnitOnly', width: 80, editableCellTemplate: 'oneUnitOnlyTemplate.html'},
           	{field:'FFOnly', width: 80, editableCellTemplate: 'FFOnlyTemplate.html'}
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: true,
        enableSorting: true,
        showColumnMenu: true,
        showFilter: true,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 120}, 
           	{field:'Descr', displayName:'Description',width: 300},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("BonusImpulseUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}


	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		console.log(row);
		targetID = row["@id"];
		console.log("Looking for ID ",targetID);
		gotSome = false;
		angular.forEach($scope.Data, function(v,i){
			if (v["@id"] == targetID) {
				//console.log("The update is on the first grid");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
				gotSome = true;
			}
		});
		if (!gotSome) { 
				angular.forEach($scope.ModData, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the mod data grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
					gotSome = true;
				}
			})
		}
		if (!gotSome) {
			if ('Code' in row) {
				//console.log("The update is on the mod data grid because it has a property called Code");
				DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
			} else {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Score: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.ModData.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		//console.log("Change Data Callback",d);
		$scope.$apply();
	}

	$scope.changeModData = function(d) {
		//console.log("Change Mod Data Callback",d)
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.ModEntity, Data: $scope.ModData, Callback: $scope.changeModData}
	]);
	
}])
;