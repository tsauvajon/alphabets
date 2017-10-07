

/**
 * Sample transaction processor function.
 */



function onSetSensorTemp(setSensorTemp) {
    setSensorTemp.gauge.sensorTemp = setSensorTemp.newSensorValue; 
    return getAssetRegistry('org.acme.sample.Sensor')
      .then(function (assetRegistry) {
          return assetRegistry.update(setSensorTemp.gauge);
      });
}

function onChangeThermostatTemp(changeThermostat) {
  var diff = Math.abs(changeThermostat.thermostat.sensorTemp - changeThermostat.newThermostatValue);
    if (diff <= 3) {
      changeThermostat.thermostat.thermostatTemp = changeThermostat.newThermostatValue;
      return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(changeThermostat.thermostat);
      });
    } else {
      //reject transaction
      throw new Error("Too much difference! Current sensor reading " + changeThermostat.thermostat.sensorTemp);
    }
}

function onCompareWeather(compareWeather) {
  //Make life easier. Put all values for this function in vars.
  var outsideTemp = compareWeather.outsideTemp;
  var feelsLike = compareWeather.feelsLike;
  var thermostatTemp = compareWeather.recommend.thermostatTemp;
  
  if (outsideTemp == feelsLike){
     //If the temps are the same then create req's
    
    //It's HOT
    if (outsideTemp >= 26) {
      if (thermostatTemp != 22) {
        compareWeather.recommend.recommendation = "Boy! It is HOT! The recommended thermostat " +
          "setting is 22 C. The thermostat is being adjusted from " + thermostatTemp + ".";
        compareWeather.sensor.thermostatTemp = 22;
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
      });
      } else {
        compareWeather.recommend.recommendation = "Boy! It is HOT! The recommended thermostat " +
          "setting is 22 C. Way to go! Your thermostat is already optimally set.";
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
        });
      }
    //Temperate weather  
    } else if (outsideTemp >= 20 && outsideTemp < 26) {
      if (thermostatTemp != 20) {
        compareWeather.recommend.recommendation = "Nice weather you're having! The recommended" 
          + " thermostat setting is 20 C. The thermostat is being adjusted from " + thermostatTemp +
          ".";
        compareWeather.sensor.thermostatTemp = 20;
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
      });
      } else {
        compareWeather.recomend.recommendation = "Great weather! The recommended thermostat " +
          "setting is 20 C.Way to go! Your thermostat is already optimally set.";
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
        });
      }
    //Cooler temps
    } else {
      if (thermostatTemp != 19) {
        compareWeather.recommend.recommendation = "Getting chilly! The recommended" 
          + " thermostat setting is 19 C. The thermostat is being adjusted from " + thermostatTemp +
          ".";
        compareWeather.recommend.thermostatTemp = 19;
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
      });
      } else {
        compareWeather.recommend.recommendation = "It's getting chilly! The recommended thermostat " +
          "setting is 19 C.Way to go! Your thermostat is already optimally set.";
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
        });
      }
    }
  }
  else {
    //Set the req's off of what it feelsLike and not what the actual temp is
    
    //It's HOT
    if (feelsLike >= 26) {
      if (thermostatTemp != 22) {
        compareWeather.recommend.recommendation = "Boy! It feels HOT! The recommended thermostat " +
          "setting is 22 C. The thermostat is being adjusted from " + thermostatTemp + ".";
        compareWeather.recommend.thermostatTemp = 22;
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
      });
      } else {
        compareWeather.recommend.recommendation = "Boy! It feels HOT! The recommended thermostat " +
          "setting is 22 C. Way to go! Your thermostat is already optimally set.";
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
        });
      }
    //Temperate weather  
    } else if (feelsLike >= 20 && feelsLike < 26) {
      if (thermostatTemp != 20) {
        compareWeather.recommend.recommendation = "It feels quite nice! The recommended" 
          + " thermostat setting is 20 C. The thermostat is being adjusted from " + thermostatTemp +
          ".";
        compareWeather.recommend.thermostatTemp = 20;
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
      });
      } else {
        compareWeather.recommend.recommendation = "It feels nice out! The recommended thermostat " +
          "setting is 20 C.Way to go! Your thermostat is already optimally set.";
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
        });
      }
    //Cooler temps
    } else {
      if (feelsLike != 19) {
        compareWeather.recommend.recommendation = "Brr! Where is my jacket? The recommended" 
          + " thermostat setting is 19 C. The thermostat is being adjusted from " + thermostatTemp +
          ".";
        compareWeather.recommend.thermostatTemp = 19;
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
      });
      } else {
        compareWeather.recommend.recommendation = "Brr! Where is the heat? The recommended thermostat "
          + "setting is 19 C.Way to go! Your thermostat is already optimally set.";
        return getAssetRegistry('org.acme.sample.Sensor')
        .then(function (assetRegistry) {
          return assetRegistry.update(compareWeather.recommend);
        });
      }
    }
  }
}
