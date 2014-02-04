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
	$scope.title = 'Unit Types'

	UnitTypeSocket.connect();
	//UnitTypeSocket.send('Get')

	$scope.gridOptions = { 
		data: 'UnitTypes',
		enableCellSelection: true,
        enableRowSelection: false,
        enableCellEdit: true,

        columnDefs: [
        	{field:'Name'}, 
        	{field:'Rating'},
        	{field:'Men'},
        	{field:'Size'},
        	{field:'Firepower'},
        	{field:'DrillBook'}
        ]
	};

	UnitTypeSocket.subscribe (function(e) {
		$scope.UnitTypes = JSON.parse(e);
		$scope.$apply();
	});
	
}]);
