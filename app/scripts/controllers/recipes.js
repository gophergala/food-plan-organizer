'use strict';

/**
 * @ngdoc function
 * @name foodPlanOrganizerApp.controller:MainCtrl
 * @description
 * # RecipesCtrl
 * Controller of the foodPlanOrganizerApp
 */
angular.module('foodPlanOrganizerApp')
.factory('Recipe', ['$resource', function($resource) {
    return $resource('http://localhost:8080/recipes/', {
      id: '@id'
    }, {
      update: {
        method: 'PUT'
      }
    });
  }])
.controller('RecipesCtrl', function($scope, Recipe) {
  $scope.recipes = Recipe.query();

  $scope.deleteRecipe = function(recipeId) {
    Recipe.delete({
      id: recipeId
    }, function() {
      $scope.recipes = Recipe.query();
    });
  };
})
.controller('EditRecipeCtrl', function($scope, $routeParams, $location, Recipe) {
  $scope.recipe = Recipe.get({
    id: $routeParams.id
  });
  $scope.submit = function() {
    Recipe.update($scope.recipe, function() {
      $location.path('/recipes');
    });
  };
})
.controller('NewRecipeCtrl', function($scope, $location, Recipe) {
  $scope.recipe = {};
  $scope.submit = function() {
    Recipe.save($scope.recipe, function() {
      $location.path('/recipes');
    });
  };
});