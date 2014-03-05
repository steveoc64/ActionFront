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
		.state('cc', {
 			url: '/cc',
 			templateUrl: 'commandControl.html',
 			controller: 'CommandControlCtrl'
 		})
 		.state('cc.initTables', {
 			url: '/initTables',
 			templateUrl: 'initTables.html',
 			controller: 'InitTablesCtrl'
 		})
 		.state('cc.corpsOrders', {
 			url: '/corpsOrders',
 			templateUrl: 'corpsOrders.html',
 			controller: 'CorpsOrdersCtrl'
 		})
 		.state('cc.meOrders', {
 			url: '/meOrders',
 			templateUrl: 'meOrders.html',
 			controller: 'MEOrdersCtrl'
 		})
		.state('cc.orderArrival', {
 			url: '/orderArrival',
 			templateUrl: 'orderArrival.html',
 			controller: 'OrderArrivalCtrl'
 		})
 		.state('cc.orderActivation', {
 			url: '/orderActivation',
 			templateUrl: 'orderActivation.html',
 			controller: 'OrderActivationCtrl'
 		})
 		.state('cc.commanderAction', {
 			url: '/commanderAction',
 			templateUrl: 'commanderAction.html',
 			controller: 'CommanderActionCtrl'
 		})
		.state('mf', {
 			url: '/mf',
 			templateUrl: 'moraleFatigue.html',
 			controller: 'MoraleFatigueCtrl'
 		})
 		.state('mf.MEMorale', {
 			url: '/MEMorale',
 			templateUrl: 'MEMorale.html',
 			controller: 'MEMoraleCtrl'
 		})
 		.state('mf.MEPanic', {
 			url: '/MEPanic',
 			templateUrl: 'MEPanic.html',
 			controller: 'MEPanicCtrl'
 		})
 		.state('mf.UnitMorale', {
 			url: '/UnitMorale',
 			templateUrl: 'UnitMorale.html',
 			controller: 'UnitMoraleCtrl'
 		})
 		.state('mf.FireDisc', {
 			url: '/FireDisc',
 			templateUrl: 'FireDisc.html',
 			controller: 'FireDiscCtrl'
 		})
 		.state('mf.InitBadMorale', {
 			url: '/InitBadMorale',
 			templateUrl: 'InitBadMorale.html',
 			controller: 'InitBadMoraleCtrl'
 		})
 		.state('mf.BonusImpulse', {
 			url: '/BonusImpulse',
 			templateUrl: 'BonusImpulse.html',
 			controller: 'BonusImpulseCtrl'
 		})
 		.state('mf.MEFatigue', {
 			url: '/MEFatigue',
 			templateUrl: 'MEFatigue.html',
 			controller: 'MEFatigueCtrl'
 		})
 		.state('mf.FatigueRecovery', {
 			url: '/FatigueRecovery',
 			templateUrl: 'FatigueRecovery.html',
 			controller: 'FatigueRecoveryCtrl'
 		})
 		.state('mf.MoraleRecovery', {
 			url: '/MoraleRecovery',
 			templateUrl: 'MoraleRecovery.html',
 			controller: 'MoraleRecoveryCtrl'
 		})
		.state('mv', {
 			url: '/mv',
 			templateUrl: 'movement.html',
 			controller: 'MovementCtrl'
 		})		
 		.state('mv.GTMovement', {
 			url: '/GTMovement',
 			templateUrl: 'GTMovement.html',
 			controller: 'GTMovementCtrl'
 		})
 		.state('mv.Deployment', {
 			url: '/Deployment',
 			templateUrl: 'Deployment.html',
 			controller: 'DeploymentCtrl'
 		})
 		.state('mv.TacMovement', {
 			url: '/TacMovement',
 			templateUrl: 'TacMovement.html',
 			controller: 'TacMovementCtrl'
 		})
 		.state('mv.ArtyMovement', {
 			url: '/ArtyMovement',
 			templateUrl: 'ArtyMovement.html',
 			controller: 'ArtyMovementCtrl'
 		})
 		.state('mv.ArtyExtra', {
 			url: '/ArtyExtra',
 			templateUrl: 'ArtyExtra.html',
 			controller: 'ArtyExtraCtrl'
 		})
 		.state('mv.SKRelocate', {
 			url: '/SKRelocate',
 			templateUrl: 'SKRelocate.html',
 			controller: 'SKRelocateCtrl'
 		}) 		
 		.state('mv.BUAMovement', {
 			url: '/BUAMovement',
 			templateUrl: 'BUAMovement.html',
 			controller: 'BUAMovementCtrl'
 		})
 		.state('mv.FormationChange', {
 			url: '/FormationChange',
 			templateUrl: 'FormationChange.html',
 			controller: 'FormationChangeCtrl'
 		})
		.state('fire', {
 			url: '/fire',
 			templateUrl: 'fire.html',
 			controller: 'FireCtrl'
 		})		
 		.state('fire.SKFire', {
 			url: '/SKFire',
 			templateUrl: 'SKFire.html',
 			controller: 'SKFireCtrl'
 		})
 		.state('fire.VolleyFire', {
 			url: '/VolleyFire',
 			templateUrl: 'VolleyFire.html',
 			controller: 'VolleyFireCtrl'
 		})
 		.state('fire.FireFight', {
 			url: '/FireFight',
 			templateUrl: 'FireFight.html',
 			controller: 'FireFightCtrl'
 		})
 		.state('fire.ArtyFire', {
 			url: '/ArtyFire',
 			templateUrl: 'ArtyFire.html',
 			controller: 'ArtyFireCtrl'
 		})
 		.state('fire.Bouncethru', {
 			url: '/Bouncethru',
 			templateUrl: 'Bouncethru.html',
 			controller: 'BouncethruCtrl'
 		})
 		.state('fire.CounterBtyFire', {
 			url: '/CounterBtyFire',
 			templateUrl: 'CounterBtyFire.html',
 			controller: 'CounterBtyFireCtrl'
 		})
 		.state('fire.BridgeFire', {
 			url: '/BridgeFire',
 			templateUrl: 'BridgeFire.html',
 			controller: 'BridgeFireCtrl'
 		})
		.state('ca', {
 			url: '/ca',
 			templateUrl: 'ca.html',
 			controller: 'CaCtrl'
 		})		 		
 		.state('ca.DefFire', {
 			url: '/DefFire',
 			templateUrl: 'DefFire.html',
 			controller: 'DefFireCtrl'
 		})
 		.state('ca.FormSquare', {
 			url: '/FormSquare',
 			templateUrl: 'FormSquare.html',
 			controller: 'FormSquareCtrl'
 		})
 		.state('ca.LimberIfCharged', {
 			url: '/LimberIfCharged',
 			templateUrl: 'LimberIfCharged.html',
 			controller: 'LimberIfChargedCtrl'
 		})
 		.state('ca.ShockValue', {
 			url: '/ShockValue',
 			templateUrl: 'ShockValue.html',
 			controller: 'ShockValueCtrl'
 		})
 		.state('ca.CACav', {
 			url: '/CACav',
 			templateUrl: 'CACav.html',
 			controller: 'CACavCtrl'
 		})
 		.state('ca.CAInf', {
 			url: '/CAInf',
 			templateUrl: 'CAInf.html',
 			controller: 'CAInfCtrl'
 		})
 		.state('ca.CAResult', {
 			url: '/CAResult',
 			templateUrl: 'CAResult.html',
 			controller: 'CAResultCtrl'
 		}) 		
 		;
 }])
.factory('DataSocket', ["$rootScope", "$state", "$location", "$window", function($rootScope, $state, $location, $window) {
  var service = {};
  $rootScope.FilterValues = {};

  service.connect = function(subscriptions) {
  	// subscriptions is an array of objects in the form
  	// Entity, Data, Callback
  	service.subscriptions = subscriptions;

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
		if (service.reconnecting === true) {
			console.log("Server is back up - Forcing page reload");
			$window.location.reload();
		} else {
	    	angular.forEach(service.subscriptions, function(v,i){
	    		initMsg = JSON.stringify({"Action":"List", "Entity":v.Entity});
	    		service.ws.send(initMsg);
	    	});
			service.isOpen = true;
		}
	}

	ws.onclose = function(e) {
		console.log("Reconnecting with server");
		service.isOpen = false;
		service.reconnecting = true;
		service.ws = null;

		var timeout = setTimeout(function() {
	  		service.connect(service.subscriptions);
		},2000);
	}

	ws.onerror = function(e) {
	}

    service.ws = ws;
  }
 
  service.send = function(message) {
  	if (!service.isOpen) {
  		// if we are dead, try to reopen the connection on demand
  		//service.connect(service.subscriptions);
  		console.log("Cannot send, server is still down ....");
  	} else {
    	service.ws.send(message); 		
  	}
  }
 
  return service;
}])
.controller("UnitTypesCtrl", ["$scope", "$rootScope", function($scope, $rootScope){
	$scope.FilterUpdate = function() {
		$rootScope.FilterValues = {"Nation":$scope.Nation, "Year":$scope.Year, "Name":$scope.Name};
		$scope.$broadcast('FilterUpdate', $rootScope.FilterValues);
	}
}])
.controller("CommandControlCtrl", ["$scope", "$rootScope", function($scope, $rootScope){
}])
.controller("MoraleFatigueCtrl", ["$scope", "$rootScope", function($scope, $rootScope){
}])
.controller("MovementCtrl", ["$scope", "$rootScope", function($scope, $rootScope){
}])
.controller("FireCtrl", ["$scope", "$rootScope", function($scope, $rootScope){
}])
.controller("CaCtrl", ["$scope", "$rootScope", function($scope, $rootScope){
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);			
			$scope.update(evt.targetScope.row);
		} 
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);		
			$scope.update(evt.targetScope.row);
		} 
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);			
			$scope.update(evt.targetScope.row);
		} 
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);
			$scope.update(evt.targetScope.row);
		} 
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);
			$scope.update(evt.targetScope.row);
		} 
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);
			$scope.update(evt.targetScope.row);
		} 
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
	$scope.title = "Small Arms";
	$scope.docs = "";
	$scope.Entity = "Equip";
	$scope.Data2 = [];
	$scope.title2 = "Artillery Ranges (Grids)";
	$scope.docs2 = "Table 17.2";
	$scope.Entity2 = "ArtRange";

	//DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields: ['Name'],
        	directions: ['asc']
        },
        columnDefs: [
           	{field:'Name', width: 180}, 
           	{field:'SK', displayName: 'Skirmish', width: 100}, 
           	{field:'Volley', width: 100}, 
           	{field:'Close', width: 100}, 
           	{field:'Long', width: 100}, 
        ]
	};
	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields: ['Long'],
        	directions: ['asc']
        },
        columnDefs: [
           	{field:'Weight', width: 120}, 
           	{field:'Short', width: 80}, 
           	{field:'Medium', width: 80}, 
           	{field:'Long', width: 80}, 
        ]
	};

	$scope.update = function(row) {
		console.log("EquipUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {													
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
					gotSome = true;
				}
			});
			if (!gotSome) { 
					angular.forEach($scope.Data2, function(v,i){
					if (v["@id"] == targetID) {
						//console.log("The update is on the mod data grid");
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) {
				if ('Short' in row) {
					//console.log("The update is on the mod data grid because it has a property called Code");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Name: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', Weight: '~ ??? ~'})
    }

    $scope.changeData = function(d) {
    	$scope.$apply();
    }

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData}
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);
			$scope.update(evt.targetScope.row);
		} 
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);
			$scope.update(evt.targetScope.row);
		} 
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);
			$scope.update(evt.targetScope.row);
		} 
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
false
	//DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);
			$scope.update(evt.targetScope.row);
		} 
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Dice'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Dice', width: 80}, 
           	{field:'Points', displayName: 'Activation Points', width: 200}
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Points'],
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {													
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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
        showColumnMenu: false,
        showFilter: false,
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
	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {											
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData2}
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {									
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {							
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Rating: '~ ??? ~'})
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {					
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Rating: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.ModData.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		console.log("Change Data Callback",d);
		$scope.$apply();
	}

	$scope.changeModData = function(d) {
		console.log("Change Mod Data Callback",d)
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
	$scope.moddocs = "Table 13.3A";
	$scope.Entity = "FireDisciplineTest";
	$scope.ModEntity = "FireDisciplineMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {				
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Rating: '~ ??? ~'})
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
	$scope.moddocs = "Table 22.2A";
	$scope.Entity = "InitialBadMorale";
	$scope.ModEntity = "InitialBadMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 220},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("InitBadUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {		
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Score', width: 60}, 
           	{field:'Descr', displayName:'Description',width: 280}, 
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
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 300},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("BonusImpulseUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {	
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
.controller("MEFatigueCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.maintitle = "ME Fatigue Test";
	$scope.modtitle = "ME Fatigue Test Modifiers";
	$scope.docs = "Table 22.1";
	$scope.moddocs = "Table 22.1A";
	$scope.Entity = "MEFatigue";
	$scope.ModEntity = "MEFatigueMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Score', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 300}, 
           	{field:'OnlyIfNotLastTurn', displayName:'Only if not last turn',width: 120, editableCellTemplate: 'onlyNotLastTurnTemplate.html'},
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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
		console.log("MEFatigueUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
.controller("FatigueRecoveryCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.maintitle = "Fatigue Recovery";
	$scope.modtitle = "Fatigue Recovery Modifiers";
	$scope.docs = "Table 22.3";
	$scope.moddocs = "Table 22.2B";
	$scope.Entity = "FatigueRecovery";
	$scope.ModEntity = "FatigueRecoveryMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Score', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 300}, 
           	{field:'Recover', displayName:'Recover Levels',width: 120},
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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
		console.log("FatigueRecoveryUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
.controller("MoraleRecoveryCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.maintitle = "Bad Morale Recovery";
	$scope.modtitle = "Bad Morale Recovery Modifiers";
	$scope.docs = "Table 22.4";
	$scope.moddocs = "Table 22.4";
	$scope.Entity = "BadMoraleRec";
	$scope.ModEntity = "BadMoraleRecMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['GoodMorale'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Rating', width: 160}, 
           	{field:'GoodMorale', displayName:'Good Morale',width: 120}, 
           	{field:'TryAgain', displayName:'Try Again',width: 120},
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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
		console.log("MoraleRecoveryUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Rating: '~ ??? ~'})
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
.controller("GTMovementCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.title = "Grand Tactical Movement";
	$scope.docs = "Table 9.3";
	$scope.Entity = "GTMove";

	//DataSocket.connect($scope);

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
           	fields: ['METype'],
        	directions: ['asc']    	
        },
        columnDefs: [
           	{field:'METype', width: 240}, 
           	{field:'D1', displayName: 'Deployed',width: 120},
           	{field:'D2', displayName: 'Bde Out',width: 120},
           	{field:'D3', displayName: 'Deploying',width: 120},
           	{field:'D4', displayName: 'Cond Column',width: 120},
           	{field:'D5', displayName: 'Reg Column',width: 120},
           	{field:'D6', displayName: 'Extd Column',width: 120}
        ]
	};

	$scope.update = function(row) {
		console.log("GTMoveUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	$scope.updateFilters = function() {
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		if (evt.targetScope.row.entity[evt.targetScope.col.field] != $scope.saveCell) {
			console.log('Was:',$scope.saveCell);
			$scope.update(evt.targetScope.row);
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', METype: '~ ??? ~'})
    }

    $scope.changeData = function(d) {
    	console.log("New data set ",d);
    	$scope.$apply(function() {
    		$scope.Data = d;
    	});
    }

	DataSocket.connect([
		{"Entity": $scope.Entity, "Data": $scope.Data, "Callback": $scope.changeData}
	]);
	
}])
.controller("DeploymentCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.ModData = [];
	$scope.DepData = [];
	$scope.maintitle = "Deployment Result";
	$scope.modtitle = "Deployment Modifiers";
	$scope.deptitle = "Deployment States";
	$scope.docs = "Table 9.1";
	$scope.moddocs = "Table 9.1B";
	$scope.depdocs = "Table 9.2";
	$scope.Entity = "Deployment";
	$scope.ModEntity = "DeploymentMod";
	$scope.DepEntity = "DeploymentState";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Score', width: 80}, 
           	{field:'Change', displayName:'Deployment States',width: 200}, 
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'ModData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
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

	$scope.gridOptionsDeps = { 
		data: 'DepData',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['State'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'State', width: 120}, 
           	{field:'Descr', displayName:'Description',width: 200},
           	{field:'Notes', width: 400},
           	{field:'ContactFront', width: 200},
           	{field:'ContactFlank', width: 200},
           	{field:'ContactShaken', width: 80},
        ]
	};

	$scope.update = function(row) {
		console.log("DeploymentUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
					angular.forEach($scope.DepData, function(v,i){
					if (v["@id"] == targetID) {
						//console.log("The update is on the mod data grid");
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.DepEntity,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) {
				if ('Code' in row) {
					//console.log("The update is on the mod data grid because it has a property called Code");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.ModEntity,"Data":row}));
				} else if ('Notes' in row) {
					//console.log("The update is on the mod data grid because it has a property called Code");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.DepEntity,"Data":row}));

				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Rating: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.ModData.push({"@id": '0', Code: '~ ??? ~'})
    }
    $scope.newRow3 = function() {
    	$scope.DepData.push({"@id": '0', Code: '~ ??? ~'})
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
		{Entity: $scope.ModEntity, Data: $scope.ModData, Callback: $scope.changeData},
		{Entity: $scope.DepEntity, Data: $scope.DepData, Callback: $scope.changeData}
	]);
	
}])
.controller("TacMovementCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.maintitle = "Tactical Movement";
	$scope.modtitle = "Extra Move";
	$scope.docs = "Table 14.1, 14.7";
	$scope.moddocs = "Table 14.3";
	$scope.Entity = "TacMove";
	$scope.Entity2 = "AdditionalMove";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Move'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'UnitType', width: 100}, 
           	{field:'Move', width: 80}, 
           	{field:'LtWood' ,width: 80}, 
           	{field:'HvWood' ,width: 80}, 
           	{field:'Mud', width: 80}, 
           	{field:'Marsh',width: 80}, 
           	{field:'LoWall',width: 80}, 
           	{field:'HiWall',width: 80}, 
        ]
	};

	$scope.gridOptionsMods = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Terrain'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Terrain', width: 120}, 
           	{field:'D1', displayName:'Extra 1',width: 60},
           	{field:'D2', displayName:'Extra 2',width: 60},
           	{field:'D3', displayName:'Extra 3',width: 60}
        ]
	};

	$scope.update = function(row) {
		console.log("TacMoveUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
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
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) {
				if ('Terrain' in row) {
					//console.log("The update is on the mod data grid because it has a property called Code");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Rating: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData}
	]);
	
}])
.controller("ArtyMovementCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.Data3 = [];
	$scope.Data4 = [];
	$scope.title = "Artillery Movement";
	$scope.title2 = "Horse Loss";
	$scope.title3 = "Artillery Relocation";
	$scope.title4 = "Relocation Mods";
	$scope.docs = "Table 14.11";
	$scope.docs2 = "Table 14.13";
	$scope.docs3 = "Table 14.14";
	$scope.docs4 = "Table 14.14A"
	$scope.Entity = "ArtyMove";
	$scope.Entity2 = "ArtyHorseLoss";
	$scope.Entity3 = "ArtyRelocate";
	$scope.Entity4 = "ArtyRelocateMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Regular'],
        	directions:['desc']
        },
	    groups: ['Class'],
        groupsCollapsedByDefault: false,
        columnDefs: [
           	{field:'Class', visible: false, width: 80}, 
           	{field:'Guns', width: 80}, 
           	{field:'Regular', width: 80}, 
           	{field:'Gallop', width: 80}, 
           	{field:'Prolong', width: 80}, 
        ]
	};

	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Terrain'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Terrain', width: 120}, 
           	{field:'Loss', width: 80},
        ]
	};

	$scope.gridOptions3 = { 
		data: 'Data3',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['R6'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Class', width: 120}, 
           	{field:'R6', displayName:'Relocate 6h',width: 80},
           	{field:'R5', displayName:'5h',width: 80},
           	{field:'R4', displayName:'4h',width: 80},
           	{field:'R3', displayName:'3h',width: 80},
           	{field:'R2', displayName:'2h',width: 80},
           	{field:'R1', displayName:'1h',width: 80},
           	{field:'R0', displayName:'0h',width: 80},
           	{field:'W6', displayName:'Withdraw 6h',width: 80},
           	{field:'W5', displayName:'5h',width: 80},
           	{field:'W4', displayName:'4h',width: 80},
           	{field:'W3', displayName:'3h',width: 80},
           	{field:'W2', displayName:'2h',width: 80},
           	{field:'W1', displayName:'1h',width: 80},
           	{field:'W0', displayName:'0h',width: 80},
        ]
	};
	$scope.gridOptions4 = { 
		data: 'Data4',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate4.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 180},
           	{field:'Value', width: 80},
        ]
	};

	$scope.update = function(row) {
		console.log("ArtyMovementUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
					angular.forEach($scope.Data3, function(v,i){
					if (v["@id"] == targetID) {
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) { 
					angular.forEach($scope.Data4, function(v,i){
					if (v["@id"] == targetID) {
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity4,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) {
				if ('Terrain' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else if ('R6' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
				} else if ('Code' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity4,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Class: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', Terrain: '~ ??? ~'})
    }
    $scope.newRow3 = function() {
    	$scope.Data3.push({"@id": '0', Class: '~ ??? ~'})
    }
    $scope.newRow4 = function() {
    	$scope.Data4.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
		{Entity: $scope.Entity3, Data: $scope.Data3, Callback: $scope.changeData},
		{Entity: $scope.Entity4, Data: $scope.Data4, Callback: $scope.changeData},
	]);
	
}])
.controller("ArtyExtraCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.Data3 = [];
	$scope.title = "Double Team Mods";
	$scope.title2 = "Recover Guns";
	$scope.title3 = "Recovery Mods";
	$scope.docs = "Table 17.9 (Req 11+ to DblTeam)";
	$scope.docs2 = "Table 17.10";
	$scope.docs3 = "Table 17.10A";
	$scope.Entity = "DoubleTeamMod";
	$scope.Entity2 = "ArtFate";
	$scope.Entity3 = "ArtFateMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['desc']
        },
        columnDefs: [
           	{field:'Code', width: 120}, 
           	{field:'Value', width: 80}, 
        ]
	};

	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Situation', width: 200}, 
           	{field:'Score', width: 80},
        ]
	};

	$scope.gridOptions3 = { 
		data: 'Data3',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 180},
           	{field:'Value', width: 80},
        ]
	};

	$scope.update = function(row) {
		console.log("ArtyExtraUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
					angular.forEach($scope.Data3, function(v,i){
					if (v["@id"] == targetID) {
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) {
				if ('Situation' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
				} else if ('Descr' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Code: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', Situation: '~ ??? ~'})
    }
    $scope.newRow3 = function() {
    	$scope.Data3.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
		{Entity: $scope.Entity3, Data: $scope.Data3, Callback: $scope.changeData},
	]);
	
}])
.controller("BUAMovementCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.title = "Occupy / Exit Structure";
	$scope.title2 = "Modifiers";
	$scope.docs = "Table 14.6, 14.7";
	$scope.docs2 = "Table 14.7A";
	$scope.Entity = "BUAMove";
	$scope.Entity2 = "BUAMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Exit'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Rating', width: 120}, 
           	{field:'Ordered', width: 80}, 
           	{field:'Exit', width: 80}, 
        ]
	};

	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 60}, 
           	{field:'Descr', displayName:'Description',width: 250},
           	{field:'Value',width:80}
        ]
	};

	$scope.update = function(row) {
		console.log("BUAMovementUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
				if ('Code' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Rating: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
	]);
	
}])
.controller("SKRelocateCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.Data3 = [];
	$scope.title = "Skirmisher Relocate";
	$scope.title2 = "Skirmisher Relocate Mods";
	$scope.title3 = "Support Distance";
	$scope.docs = "Table 14.9";
	$scope.docs2 = "Table 14.9A";
	$scope.docs3 = "Table 14.10";
	$scope.Entity = "SKRelocate";
	$scope.Entity2 = "SKRelocateMod";
	$scope.Entity3 = "SKSupport";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Retire'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Rating', width: 100}, 
           	{field:'Retire', width: 60}, 
           	{field:'Move', width: 60}, 
           	{field:'Bold', width: 60}, 
        ]
	};

	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 60}, 
           	{field:'Descr', displayName:'Description',width: 200},
           	{field:'Value',width:60}
        ]
	};
	$scope.gridOptions3 = { 
		data: 'Data3',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Mode'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Mode', width: 80}, 
           	{field:'Marchfeld', width: 100},
           	{field:'Rolling',width:60},
           	{field:'Rough',width:60}
        ]
	};

	$scope.update = function(row) {
		console.log("SKRelocateUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
					angular.forEach($scope.Data3, function(v,i){
					if (v["@id"] == targetID) {
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) {
				if ('Code' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else if ('Mode' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Rating: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
		{Entity: $scope.Entity3, Data: $scope.Data3, Callback: $scope.changeData},
	]);
	
}])
.controller("FormationChangeCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.title = "Formation Change";
	$scope.docs = "Table 14.4";
	$scope.Entity = "FormationChange";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['To'],
        	directions:['asc']
        },
        groups: ['Era','From'],
        columnDefs: [
           	{field:'Era', visible: false,width: 100}, 
           	{field:'From', visible: false, width: 200}, 
           	{field:'To', width: 200}, 
           	{field:'Trained', width: 200}, 
           	{field:'Untrained', width: 200}, 
        ]
	};


	$scope.update = function(row) {
		console.log("FormationChangeUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
					gotSome = true;
				}
			});
			if (!gotSome) {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Rating: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
	]);
	
}])
.controller("SKFireCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.Data3 = [];
	$scope.Data4 = [];
	$scope.title = "Fire Chart";
	$scope.title2 = "Fire Effect";
	$scope.title3 = "Skirmish Fire Modifiers";
	$scope.title4 = "To Hit Values";
	$scope.docs = "Table 15.2";
	$scope.docs2 = "Table 15.2";
	$scope.docs3 = "Table 3.1";
	$scope.docs4 = "Table 3.2"
	$scope.Entity = "FireChart";
	$scope.Entity2 = "FireEffect";
	$scope.Entity3 = "FireSKMod";
	$scope.Entity4 = "SKEffect";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', width: 60}, 
           	{field:'SmallArms', displayName:'Skirmish %',width: 100}, 
        ]
	};

	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', width: 60}, 
           	{field:'Dice', width: 60},
           	{field:'Descr', displayName:'Description',width:120}
        ]
	};

	$scope.gridOptions3 = { 
		data: 'Data3',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 60}, 
           	{field:'Descr', displayName:'Description',width: 220},
           	{field:'Value', width: 60},
        ]
	};
	$scope.gridOptions4 = { 
		data: 'Data4',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate4.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ECode', width: 60}, 
           	{field:'Dice', width: 60},
           	{field:'Descr', displayName:'Target',width: 120},
        ]
	};

	$scope.update = function(row) {
		console.log("SKFireUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
					angular.forEach($scope.Data3, function(v,i){
					if (v["@id"] == targetID) {
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) { 
					angular.forEach($scope.Data4, function(v,i){
					if (v["@id"] == targetID) {
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity4,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) {
				if ('Dice' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else if ('ECode' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity4,"Data":row}));
				} else if ('Code' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', ID: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', ID: '~ ??? ~'})
    }
    $scope.newRow3 = function() {
    	$scope.Data3.push({"@id": '0', Code: '~ ??? ~'})
    }
    $scope.newRow4 = function() {
    	$scope.Data4.push({"@id": '0', ECode: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
		{Entity: $scope.Entity3, Data: $scope.Data3, Callback: $scope.changeData},
		{Entity: $scope.Entity4, Data: $scope.Data4, Callback: $scope.changeData},
	]);
	
}])
.controller("VolleyFireCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.Data3 = [];
	$scope.title = "Fire Chart";
	$scope.title2 = "Fire Effect";
	$scope.title3 = "Fire Modifiers";
	$scope.docs = "Table 15.2";
	$scope.docs2 = "Table 15.2";
	$scope.docs3 = "Table 15.2A & B";
	$scope.Entity = "FireChart";
	$scope.Entity2 = "FireEffect";
	$scope.Entity3 = "FireMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', width: 60}, 
           	{field:'SmallArms', displayName:'Musket %',width: 100}, 
        ]
	};

	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', width: 60}, 
           	{field:'Dice', width: 60},
           	{field:'Descr', displayName:'Description',width:120}
        ]
	};

	$scope.gridOptions3 = { 
		data: 'Data3',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['desc']
        },
        columnDefs: [
           	{field:'Code', width: 120}, 
           	{field:'Descr', displayName:'Description',width: 300},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("VolleyFireUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
					angular.forEach($scope.Data3, function(v,i){
					if (v["@id"] == targetID) {
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) {
				if ('Dice' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else if ('Code' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', ID: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', ID: '~ ??? ~'})
    }
    $scope.newRow3 = function() {
    	$scope.Data3.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
		{Entity: $scope.Entity3, Data: $scope.Data3, Callback: $scope.changeData},
	]);
	
}])
.controller("FireFightCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.title = "Fire Fight Results";
	$scope.title2 = "Fire Fight Modifiers";
	$scope.docs = "Table 15.2";
	$scope.docs2 = "Table 15.2";
	$scope.Entity = "FireFight";
	$scope.Entity2 = "FireFightMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Dice'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Dice', width: 60}, 
           	{field:'Descr', displayName:'Description',width:200},
           	{field:'Fallback', width:80, editableCellTemplate: 'fallbackTemplate.html'},
           	{field:'HoldCover', width:80, editableCellTemplate: 'holdTemplate.html'},
           	{field:'Disorder', width:80, editableCellTemplate: 'disorderTemplate.html'},
           	{field:'Rout', width:80, editableCellTemplate: 'routTemplate.html'},
        ]
	};

	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['desc']
        },
        columnDefs: [
           	{field:'Code', width: 120}, 
           	{field:'Descr', displayName:'Description',width: 300},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("FireFightUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
				if ('Code' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', ID: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', ID: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
	]);
	
}])
.controller("ArtyFireCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.Data3 = [];
	$scope.title = "Fire Effect";
	$scope.title2 = "Fire Chart";
	$scope.title3 = "Fire Modifiers";
	$scope.docs = "Table 15.2";
	$scope.docs2 = "Table 17.1";
	$scope.docs3 = "Table 17.2A";
	$scope.Entity = "FireEffect";
	$scope.Entity2 = "FireChart";
	$scope.Entity3 = "ArtMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', width: 60}, 
           	{field:'Dice', width: 60},
           	{field:'Descr', displayName:'Description',width:120}
        ]
	};

	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', width: 60}, 
           	{field:'LtArt', displayName:'Light',width: 60}, 
           	{field:'MdArt', displayName:'Medium',width: 60}, 
           	{field:'MdHvArt', displayName:'MdHeavy',width: 60}, 
           	{field:'HvArt', displayName:'Heavy',width: 60}, 
        ]
	};


	$scope.gridOptions3 = { 
		data: 'Data3',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['desc']
        },
        columnDefs: [
           	{field:'Code', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 280},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("ArtyFireUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
					angular.forEach($scope.Data3, function(v,i){
					if (v["@id"] == targetID) {
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) {
				if ('Dice' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
				} else if ('LtArt' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else if ('Code' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', ID: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', ID: '~ ??? ~'})
    }
    $scope.newRow3 = function() {
    	$scope.Data3.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
		{Entity: $scope.Entity3, Data: $scope.Data3, Callback: $scope.changeData},
	]);
	
}])
.controller("BouncethruCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.title = "Bounce Through Effect";
	$scope.title2 = "Bounce Through Mods";
	$scope.docs = "Table 18.3";
	$scope.docs2 = "Table 18.3";
	$scope.Entity = "Bouncethru";
	$scope.Entity2 = "BouncethruMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Score', width: 60}, 
           	{field:'Light', width: 120},
           	{field:'Medium', width:120},
           	{field:'MediumHv', width:120},
           	{field:'Heavy', width:120},
        ]
	};

	$scope.gridOptions2= { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['desc']
        },
        columnDefs: [
           	{field:'Code', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 280},
           	{field:'Value', width: 60},
        ]
	};

	$scope.update = function(row) {
		console.log("BounceThruUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
				if ('Code' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Score: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	$scope.Data2.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
	]);
	
}])
.controller("CounterBtyFireCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.title = "Counter Bty Fire Results";
	$scope.docs = "Table 17.5";
	$scope.Entity = "CounterBty";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Score', width: 100}, 
           	{field:'Horses', width: 100},
           	{field:'Crew', width: 100},
           	{field:'LHorses', displayName:'Limbered - Horses',width: 200},
           	{field:'LCrew', displayName: 'Limbered - Crew',width: 200},
           	{field:'Caisson', displayName:'Caisson Hit', editableCellTemplate: 'caissonTemplate.html',width: 100},
        ]
	};

	$scope.update = function(row) {
		console.log("CounterBtyUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
					gotSome = true;
				}
			});
			if (!gotSome) {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Score: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
	]);
}])
.controller("BridgeFireCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.title = "Set Bridges and Buildings Aflame";
	$scope.docs = "Table 17.8  (+1 for each additional Howitzer section in Bty)";
	$scope.Entity = "Aflame";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['TacMd'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Cover', width: 200}, 
           	{field:'TacMd', displayName: 'Tactical Fire < 18lb',width: 200},
           	{field:'TacHv', displayName: '18lb +', width: 100},
           	{field:'BombardMd', displayName:'Bombardment < 18lb',width: 200},
           	{field:'BombardHv', displayName: '18lb +',width: 100},
        ]
	};

	$scope.update = function(row) {
		console.log("BridgeFireUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
					gotSome = true;
				}
			});
			if (!gotSome) {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Cover: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
	]);
}])
.controller("DefFireCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.title = "Defensive Fire Effect";
	$scope.title2 = "Defensive Fire Notes";
	$scope.docs = "Table 16.1";
	$scope.docs2 = "Table 16.1A";
	$scope.Entity = "DefFire";
	$scope.Entity2 = "DefFireNote";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', visible: false,width: 60}, 
           	{field:'Target', displayName:'Attacking Unit',width: 180},
           	{field:'Hits1', displayName:'1-3 Hits', width:75},
           	{field:'Hits4', displayName:'4-5 Hits', width:75},
           	{field:'Hits6', displayName:'6-7 Hits', width:75},
           	{field:'Hits8', displayName:'8-9 Hits', width:75},
           	{field:'Hits10', displayName:'10+ Hits', width:75},
        ]
	};

	$scope.gridOptions2= { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 400},
        ]
	};

	$scope.update = function(row) {
		console.log("DefFireUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
				if ('Code' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
//    	$scope.Data.push({"@id": '0', Target: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
  //  	$scope.Data2.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
	]);
	
}])
.controller("FormSquareCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.title = "Form Square";
	$scope.title2 = "Form Square Mods";
	$scope.docs = "Table 14.5";
	$scope.docs2 = "Table 14.5A";
	$scope.Entity = "FormSquare";
	$scope.Entity2 = "FormSquareMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', visible: false,width: 1},
           	{field:'Rating', width: 100},
           	{field:'From', displayName:'From Formation',width:200},
           	{field:'Grid0', displayName:'Same Grid',width:100},
           	{field:'Grid1', displayName:'Adj Grid',width:100},
           	{field:'Grid1D', displayName:'Diagonal',width:100},
           	{field:'Grid2', displayName:'2+ Grids',width:100},
        ]
	};

	$scope.gridOptions2= { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', width: 80}, 
           	{field:'Descr', displayName:'Description',width: 200},
           	{field:'Value', width: 80},
        ]
	};

	$scope.update = function(row) {
		console.log("FormSquareUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
				if ('Code' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
//    	$scope.Data.push({"@id": '0', Target: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
  //  	$scope.Data2.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
	]);
	
}])
.controller("LimberIfChargedCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.title = "Limber Artillery if Charged";
	$scope.docs = "Table 17.4 (-1 Per Fatigue Level over Fresh)";
	$scope.Entity = "ArtLimber";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', visible: false, width: 10}, 
           	{field:'Code', width: 60}, 
           	{field:'Descr', displayName: 'Description of Threat',width: 300},
           	{field:'Score', width: 80},
           	{field:'Flee', displayName: 'Flee Distance', width: 200},
        ]
	};

	$scope.update = function(row) {
		console.log("LimberIfChargedUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
					gotSome = true;
				}
			});
			if (!gotSome) {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
	]);
}])
.controller("ShockValueCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.title = "Shock Value / ACE Calculator";
	$scope.docs = "Table 16.2";
	$scope.Entity = "ShockValue";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', displayName: 'Numeric Index', width: 120}, 
           	{field:'Label', width: 240}, 
           	{field:'Value', displayName: 'Shock Value / ACE Rating', width: 200},
        ]
	};

	$scope.update = function(row) {
		console.log("ShockValueUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row}));
					gotSome = true;
				}
			});
			if (!gotSome) {
				$scope.update(evt.targetScope.row);	
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', ID: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
	]);
}])
.controller("CACavCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.title = "Cavalry Close Action Factors";
	$scope.title2 = "General Close Action Factors";
	$scope.docs = "Table 16.2B";
	$scope.docs2 = "Table 16.2A & C";
	$scope.Entity = "CACav";
	$scope.Entity2 = "CAFactor";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', visible:false,width: 80},
           	{field:'Value', width:80},
           	{field:'Descr', displayName:'Description',width:300},
        ]
	};

	$scope.gridOptions2= { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', visible:false,width: 80}, 
           	{field:'Type', width: 120}, 
           	{field:'Value', width: 80},
           	{field:'Descr', displayName:'Description',width: 250},
        ]
	};

	$scope.update = function(row) {
		console.log("CACavUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
				if ('Type' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
//    	$scope.Data.push({"@id": '0', Target: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
  //  	$scope.Data2.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
	]);

}])
.controller("CAInfCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.title = "Infantry Close Action Factors";
	$scope.title2 = "General Close Action Factors";
	$scope.docs = "Table 16.2B";
	$scope.docs2 = "Table 16.2A & C";
	$scope.Entity = "CAInf";
	$scope.Entity2 = "CAFactor";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', visible:false,width: 80},
           	{field:'Value', width:80},
           	{field:'Descr', displayName:'Description',width:300},
        ]
	};

	$scope.gridOptions2= { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Code'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Code', visible:false,width: 80}, 
           	{field:'Type', width: 120}, 
           	{field:'Value', width: 80},
           	{field:'Descr', displayName:'Description',width: 250},
        ]
	};

	$scope.update = function(row) {
		console.log("CAInfUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
				if ('Type' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
//    	$scope.Data.push({"@id": '0', Target: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
  //  	$scope.Data2.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
	]);

}])
.controller("CAResultCtrl", ["$scope", "DataSocket", "$rootScope",function($scope, DataSocket,$rootScope){
	$scope.Data = [];
	$scope.Data2 = [];
	$scope.Data3 = [];
	$scope.title = "Close Action Results";
	$scope.title2 = "Result Code";
	$scope.title3 = "Close Action Mods";
	$scope.docs = "Table 16.3";
	$scope.docs2 = "Table 16.3";
	$scope.docs3 = "Table 16.3A";
	$scope.Entity = "CAResult";
	$scope.Entity2 = "CAResultCode";
	$scope.Entity3 = "CAResultMod";

	$scope.gridOptions = { 
		data: 'Data',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate.html',
        sortInfo: {
        	fields:['Score'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'Score', width: 60}, 
           	{field:'Descr', displayName:'Description',width:100},
           	{field:'LInf', displayName:'Lose Inf',width:80},
           	{field:'LCav', displayName:'Cavalry',width:100},
           	{field:'LArt', displayName:'Artillery',width:80},
           	{field:'VInf', displayName:'Win Inf',width:80},
           	{field:'VCav', displayName:'Cav',width:80},
        ]
	};

	$scope.gridOptions2 = { 
		data: 'Data2',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate2.html',
        sortInfo: {
        	fields:['ID'],
        	directions:['asc']
        },
        columnDefs: [
           	{field:'ID', visible: false,width: 60}, 
           	{field:'Code', displayName:'Code',width: 60}, 
           	{field:'Descr', displayName:'Description',width: 180}, 
        ]
	};

	$scope.gridOptions3 = { 
		data: 'Data3',
		enableCellSelection: true,
        enableCellEdit: true,
        enableColumnResize: true,
        enableColumnReordering: false,
        enableSorting: true,
        showColumnMenu: false,
        showFilter: false,
        showFooter: true,
        footerTemplate: 'gridFooterTemplate3.html',
        sortInfo: {
        	fields:['Value'],
        	directions:['desc']
        },
        columnDefs: [
           	{field:'Code', visible:false,width: 80}, 
           	{field:'Value', width: 60},
           	{field:'Descr', displayName:'Description',width: 280},
        ]
	};

	$scope.update = function(row) {
		console.log("CAResultUpdated -> ",row.entity);
		DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity,"Data":row.entity}));
	}

	// Capture the cell on start edit, and update if the cell contents change
	$scope.$on('ngGridEventStartCellEdit',function(evt){
		$scope.saveCell = evt.targetScope.row.entity[evt.targetScope.col.field];
	});
	$scope.$on('ngGridEventEndCellEdit', function(evt){
		// Nasty problem here - need to work out WHICH GRID this even belongs to
		row = evt.targetScope.row.entity;
		if (row[evt.targetScope.col.field] != $scope.saveCell) {
			console.log($scope.saveCell, ':', row);
			targetID = row["@id"];
			gotSome = false;
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == targetID) {
					//console.log("The update is on the first grid");
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
					angular.forEach($scope.Data3, function(v,i){
					if (v["@id"] == targetID) {
						DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
						gotSome = true;
					}
				})
			}
			if (!gotSome) {
				if ('ID' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity2,"Data":row}));
				} else if ('Value' in row) {
					DataSocket.send(JSON.stringify({"Action":"Update","Entity":$scope.Entity3,"Data":row}));
				} else {
					$scope.update(evt.targetScope.row);	
				}
			}
		}
    });

    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Score: '~ ??? ~'})
    }
    $scope.newRow2 = function() {
    	//$scope.Data2.push({"@id": '0', ID: '~ ??? ~'})
    }
    $scope.newRow3 = function() {
    	//$scope.Data3.push({"@id": '0', Code: '~ ??? ~'})
    }

	$scope.changeData = function(d) {
		$scope.$apply();
	}

	DataSocket.connect([
		{Entity: $scope.Entity, Data: $scope.Data, Callback: $scope.changeData},
		{Entity: $scope.Entity2, Data: $scope.Data2, Callback: $scope.changeData},
		{Entity: $scope.Entity3, Data: $scope.Data3, Callback: $scope.changeData},
	]);
	
}])
;


