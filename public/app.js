function socketUrl(s) {
    var l = window.location;
    return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname + (((l.port != 80) && (l.port != 443)) ? ":" + l.port : "") + l.pathname + s;
}

angular.module("unitEditor", ['ui.router', 'ngGrid'])
 .config(['$urlRouterProvider', '$stateProvider', function ($urlRouterProvider, $stateProvider) {
 	$urlRouterProvider.otherwise('/unittype');
 	$stateProvider
 		.state('unittype', {
 			url: '/unittype',
 			templateUrl: 'unittype.html',
 			controller: 'UnitTypeCtrl'
 		})
 		.state('drillbook', {
 			url: '/drillbook',
 			templateUrl: 'drillbook.html'
 		});
 }])
.factory('UnitTypeSocket', function() {
  var service = {};
 
  service.connect = function() {
    if(service.ws) { return; }

    var ws = new WebSocket(socketUrl('UnitTypeSocket'));
  
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
.controller("UnitTypeCtrl", ["$scope", "UnitTypeSocket", function($scope, UnitTypeSocket){
	$scope.UnitTypes = [];
	$scope.Ratings = ['OldGuard','Guard','Grenadier','Elite','CrackLine','Veteran','Regular','Conscript','Landwehr','Militia','Rabble'];
	$scope.DrillBooks = ['Light Infantry','French','Prussian','Russian','Austrian','British','Old School','Conscript','Militia','Mob'];
	$scope.title = 'Unit Types'

	UnitTypeSocket.connect();
	//UnitTypeSocket.send('Get')

	$scope.gridOptions = { 
		data: 'UnitTypes',
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
        sortInfo: {
        	fields: ['Name','Rating'],
        	directions: ['asc','asc']
        },

        columnDefs: [
        	{field:'Name'}, 
        	{field:'Rating', editableCellTemplate: 'ratingTemplate.html'},
        	{field:'DrillBook', editableCellTemplate: 'drillBookTemplate.html'},
        	{field:'Men'},
        	{field:'Size'},
        	{field:'Firepower'}
        ]
	};

	$scope.update = function(row) {
		//console.log(row);
		console.log("Updated -> ",row.entity);
		UnitTypeSocket.send(JSON.stringify(row.entity));
	}

/*
    $scope.$on('ngGridEventEndCellEdit', function(evt){
    	$scope.update(evt.targetScope.row);
    });
*/

	$scope.$on('ngGridEventEndCellEdit', function(evt){
		$scope.update(evt.targetScope.row);
		//console.log(evt);
		//console.log("Updated -> ",$scope.UnitTypes[evt.targetScope.row.entity]);
		//UnitTypeSocket.send(JSON.stringify($scope.UnitTypes[evt.targetEvt]));
    });


    $scope.newRow = function() {
    	$scope.UnitTypes.push({"@id": '0', Name: '~ New ~', Rating: 'Regular', Men: 720, Size: '3L', Firepower: '10', DrillBook: ''})
    }

	UnitTypeSocket.subscribe (function(e) {
		var data = JSON.parse(e);
		console.log($scope)
		console.log("New ->", data);

		if (data instanceof Array) {
			$scope.UnitTypes = data;
		} else if (data instanceof Object) {
			//console.log("Single Record", data);
			var gotSome = false;
			for (var i = 0; i < $scope.UnitTypes.length; i++) {
				if (data["@id"] === $scope.UnitTypes[i]["@id"]) {
					console.log("Updating record at pos ",i,"to",data);
					$scope.UnitTypes[i].Name = data.Name;
					$scope.UnitTypes[i].Rating = data.Rating;
					$scope.UnitTypes[i].Men = data.Men;
					$scope.UnitTypes[i].Size = data.Size;
					$scope.UnitTypes[i].Firepower = data.Firepower;
					$scope.UnitTypes[i].DrillBook = data.DrillBook;
					gotSome = true;
					$scope.$apply();
				}

			}
			if (!gotSome) {
				console.log("Add New Record");
				// First - overwrite any record in the list that has a blank ID		
				for (var i = 0; i < $scope.UnitTypes.length; i++) {
					//console.log("Pos",i,"id = ",$scope.UnitTypes[i]["@id"]);
					if ($scope.UnitTypes[i]["@id"] === "0") {
						console.log("Overwriting blank record at pos",i);
						$scope.UnitTypes[i]["@id"] = data["@id"];
						$scope.UnitTypes[i].Name = data.Name;
						$scope.UnitTypes[i].Rating = data.Rating;
						$scope.UnitTypes[i].Men = data.Men;
						$scope.UnitTypes[i].Size = data.Size;
						$scope.UnitTypes[i].Firepower = data.Firepower;
						$scope.UnitTypes[i].DrillBook = data.DrillBook;
						gotSome = true;
					}
				}

				// otherwise - just append to the list
				if (!gotSome) {
					console.log("Adding new record");
					$scope.UnitTypes.push(data);
				}
			}
		} else {
			console.log("Deleting record ",data)
			for (var i = 0; i < $scope.UnitTypes.length; i++) {
				if ($scope.UnitTypes[i]["@id"] == data) {
					console.log("delete row at index",i)
					$scope.UnitTypes.splice(i,1);
				}
			}
		}
		$scope.$apply();
	});
	
}]);
