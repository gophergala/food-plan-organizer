'use strict';
angular.module('foodPlanOrganizerApp')
.factory('Nutrient', ['$resource', function($resource) {
    return $resource('http://localhost:8080/nutrients/');
  }]);