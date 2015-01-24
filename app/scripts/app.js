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
    'ngAnimate',
    'ngMessages',
    'ngResource',
    'ngRoute',
    'ngSanitize',
    'ngTouch'
  ])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .when('/about', {
        templateUrl: 'views/about.html',
        controller: 'AboutCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  });
