function socketUrl(s) {
    var l = window.location;
    return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname + (((l.port != 80) && (l.port != 443)) ? ":" + l.port : "") + l.pathname + s;
}

angular.module("app", ['ui.router', 'ngGrid'])
 .config(['$urlRouterProvider', '$stateProvider', function ($urlRouterProvider, $stateProvider) {
 	$urlRouterProvider.otherwise('/');
 	$stateProvider
 		.state('unittypes', {
 			url: '/unittypes',
 			templateUrl: 'unittypes.html'
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
.factory('Socket', function() {
  var service = {};
 
  service.connect = function() {
    if(service.ws) { 
    	service.ws.send("init");
    	return; 
    }

    var ws = new WebSocket(socketUrl('Socket'));
  
    ws.onmessage = function(message) {
      service.callback(message.data);
    };
 
    service.ws = ws;
  }
 
  service.send = function(message) {
    service.ws.send(message);
  }
 
  service.subscribe = function(callback) {
    service.callback = callback;
  }
 
  return service;
})
.controller("InfantryCtrl", ["$scope", "Socket", function($scope, Socket){
	$scope.Data = [];
	$scope.Ratings = ['OldGuard','Guard','Grenadier','Elite','CrackLine','Veteran','Regular','Conscript','Landwehr','Militia','Rabble'];
	$scope.DrillBooks = ['Light Infantry','French','Prussian','Russian','Austrian','British','Old School','Conscript','Militia','Mob'];
	$scope.title = "L'Infanterie";

	Socket.connect();

	$scope.gridOptions = { 
		data: 'Data',
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
        	fields: ['Name'],
        	directions: ['asc']
        },

        columnDefs: [
        	{field:'Name', width: 120}, 
        	{field:'Rating', width: 100, editableCellTemplate: 'ratingTemplate.html'},
        	{field:'DrillBook', width: 100, editableCellTemplate: 'drillBookTemplate.html'},
        	{field:'Men'},
        	{field:'Size'},
        	{field:'Firepower'}
        ]
	};

	$scope.update = function(row) {
		console.log("Updated -> ",row.entity);
		Socket.send(JSON.stringify(row.entity));
	}

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
    });


    $scope.newRow = function() {
    	$scope.Data.push({"@id": '0', Name: '~ New ~', Rating: 'Regular', Men: 720, Size: '3L', Firepower: '10', DrillBook: ''})
    }

	Socket.subscribe (function(e) {
		var data = JSON.parse(e);
		console.log($scope)
		console.log("Msg ->", data);

		if (data instanceof Array) {
			// On Rx an array of data - set the whole dataset to the array
			$scope.Data = data;
		} else if (data instanceof Object) {
			// On Rx a single record 
			var gotSome = false;

			// If the ID of the record exists, update the record in the dataset
			angular.forEach($scope.Data, function(v,i){
				if (data["@id"] === v["@id"]) {
					console.log("Updating record at pos",i,"to",data);
					angular.copy(data,$scope.Data[i]);
					gotSome = true;
				}
			});

			// else if any of our records have a blank ID, overwrite that as the new record
			if (!gotSome) {
				console.log("Add New Record");
				angular.forEach($scope.Data, function(v,i){
					if (v["@id"] === "0") {
						console.log("Overwriting Blank record at pos",i);
						angular.copy(data,$scope.Data[i]);
						gotSome = true;
					}
				});
			}

			// otherwise - just append to the list
			if (!gotSome) {
				console.log("Adding new record");
				$scope.UnitTypes.push(data);
			}			
		} else {
			// On Rx a single ID 
			console.log("Deleting record ",data)
			angular.forEach($scope.Data, function(v,i){
				if (v["@id"] == data) {
					console.log("Delete row at pos",i);
					$scope.Data.splice(i,1);
				}
			});
		}

		// Sync the scope and the DOM
		$scope.$apply();
	});
	
}]);
