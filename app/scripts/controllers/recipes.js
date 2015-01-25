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
      food_id: ingredient.id,
      nutrients: []
    });
  });
  $scope.deleteIngredient = function(ingredient) {
    var index = $scope.recipe.ingredients.indexOf(ingredient);
    if (index !== -1) {
      $scope.recipe.ingredients.splice(index, 1);
    }
  };
};

var totalNutrients = function(recipe, nutrients, nutrientId) {
  var nutrient = null;
  for (var k = 0; k < nutrients.length; k++) {
    if (nutrients[k].id === nutrientId) {
      nutrient = nutrients[k];
    }
  }

  var total = 0.0;
  for (var i = 0; i < recipe.ingredients.length; i++) {
    var ingredient = recipe.ingredients[i];
    var multipler = 1;

    if (ingredient.unit === 'piece') {
      multipler = ingredient.volume;
    } else if (ingredient.unit === 'g') {
      if (nutrient.unit === 'g') {
        multipler = ingredient.volume;
      } else if (nutrient.unit === 'mg') {
        multipler = ingredient.volume / 1000.0;
      } else if (nutrient.unit === 'µg') {
        multipler = ingredient.volume / 1000.0 / 1000.0;
      } else {
        multipler = -1;
      }
    }
    // console.log(ingredient);
    for (var j = 0; j < ingredient.nutrients.length; j++) {
      var nutrient = ingredient.nutrients[j];
      if (nutrient.id === nutrientId) {
        total = total + (nutrient.nutrient_value * multipler);
      }
    }
  }
  return total;
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
.controller('RecipesCtrl', function($scope, Recipe) {
  $scope.recipes = Recipe.query();

  $scope.truncate = function(string, n) {
    return string.length > n ? string.substr(0, n - 1) + ' …' : string;
  };
  $scope.deleteRecipe = function(recipeId) {
    Recipe.delete({
      id: recipeId
    }, function() {
      $scope.recipes = Recipe.query();
    });
  };
})
.controller('EditRecipeCtrl', function($scope, $routeParams, $location, Recipe, Nutrient) {
  $scope.recipe = Recipe.get({
    id: $routeParams.id
  });
  $scope.nutrients = Nutrient.query();
  $scope.totalNutrients = totalNutrients;
  ingredientHandling($scope);

  $scope.submit = function() {
    Recipe.update($scope.recipe, function() {
      $location.path('/recipes');
    });
  };
})
.controller('NewRecipeCtrl', function($scope, $location, Recipe, Nutrient) {
  $scope.recipe = {};
  $scope.nutrients = Nutrient.query();
  $scope.totalNutrients = totalNutrients;
  ingredientHandling($scope);

  $scope.submit = function() {
    Recipe.save($scope.recipe, function() {
      $location.path('/recipes');
    });
  };
});