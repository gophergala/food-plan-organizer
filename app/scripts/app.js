'use strict';

/**
 * @ngdoc overview
 * @name foodPlanOrganizerApp
 * @description
 * # foodPlanOrganizerApp
 *
 * Main module of the application.
 */
angular
.module('foodPlanOrganizerApp', [
  'ngMessages',
  'ngResource',
  'ngRoute',
  'ngSanitize'
])
.config(function($routeProvider, $resourceProvider) {
  $resourceProvider.defaults.stripTrailingSlashes = false;

  $routeProvider
  .when('/', {
    templateUrl: 'views/calendar.html',
    controller: 'CalendarCtrl'
  })
  .when('/recipes', {
    templateUrl: 'views/recipes.html',
    controller: 'RecipesCtrl'
  })
  .when('/recipes/new', {
    templateUrl: 'views/recipes.edit.html',
    controller: 'NewRecipeCtrl'
  })
  .when('/recipes/:id/edit', {
    templateUrl: 'views/recipes.edit.html',
    controller: 'EditRecipeCtrl'
  })
  .when('/about', {
    templateUrl: 'views/about.html',
    controller: 'AboutCtrl'
  })
  .otherwise({
    redirectTo: '/'
  });
});