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
    });
  }])
.controller('RecipesCtrl', function($scope, Recipe) {
  $scope.recipes = Recipe.get();

  $scope.deleteRecipe = function(recipeId) {
    // TODO
  };
})
.controller('EditRecipeCtrl', function($scope, $routeParams, Recipe) {
  console.log('EditRecipe');
  $scope.recipe = Recipe.get({
    id: $routeParams.id
  });
  $scope.persist = function() {
    console.log('TODO UPDATE');
  };
  // TODO
})
.controller('NewRecipeCtrl', function($scope, Recipe) {
  console.log('NewRecipe');
  $scope.recipe = {};
  $scope.persist = function() {
    console.log('TODO CREATE');
  };
  // TODO
});