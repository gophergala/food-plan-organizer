'use strict';

/**
 * @ngdoc function
 * @name foodPlanOrganizerApp.controller:MainCtrl
 * @description
 * # RecipesCtrl
 * Controller of the foodPlanOrganizerApp
 */

var ingredientHandling = function ingredientHandling($scope) {
  $scope.selectedIngredient = null;
  $scope.recipe.ingredients = [];
  $scope.$watch('selectedIngredient', function(n, o) {
    if (n !== null) {
      $scope.currentIngredient = n.originalObject;
      $scope.$broadcast('add.ingredient', $scope.currentIngredient);
    }
  });
  $scope.$on('add.ingredient', function(evt, ingredient) {
    console.log(ingredient);
    if ($scope.recipe.ingredients === undefined || $scope.recipe.ingredients === null) {
      $scope.recipe.ingredients = [];
    }
    $scope.recipe.ingredients.push({
      name: ingredient.name,
      food_id: ingredient.id
    });
  });
  $scope.deleteIngredient = function(ingredient) {
    var index = $scope.recipe.ingredients.indexOf(ingredient);
    if (index !== -1) {
      $scope.recipe.ingredients.splice(index, 1);
    }
  };
};

angular.module('foodPlanOrganizerApp')
.directive('float', function() {
  return {
    require: 'ngModel',
    link: function(scope, ele, attr, ctrl) {
      ctrl.$parsers.unshift(function(viewValue) {
        return parseFloat(viewValue);
      });
    }
  };
})
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
  ingredientHandling($scope);

  $scope.submit = function() {
    Recipe.update($scope.recipe, function() {
      $location.path('/recipes');
    });
  };
})
.controller('NewRecipeCtrl', function($scope, $location, Recipe) {
  $scope.recipe = {};
  ingredientHandling($scope);

  $scope.submit = function() {
    Recipe.save($scope.recipe, function() {
      $location.path('/recipes');
    });
  };
});