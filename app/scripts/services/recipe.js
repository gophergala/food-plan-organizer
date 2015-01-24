'use strict';
angular.module('foodPlanOrganizerApp')
.factory('Recipe', ['$resource', function($resource) {
    return $resource('http://localhost:8080/recipes/', {
      id: '@id'
    }, {
      update: {
        method: 'PUT'
      }
    });
  }]);