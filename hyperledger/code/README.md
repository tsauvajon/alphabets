# Hyplerledger Composer Lab

The blockchain workshop is intended to give you a basic understanding of how a developer would interact with Hyperledger Fabric. In this workshop you will use a Browser based UI to modify chaincode, test your code and deploy your changes. You will also learn how tooling can take the code and generate API to allow for application integration through a REST-ful interface.

This lab will be broken into two parts: creating chaincode and generating API and using NodeRED to test API integration.

Your Raspberry Pi Sens HAT detects the temperature in the room or the temperature you create around it. In a real world scenario, this could be a temperature gauge in your house or in an office building. In this lab, we have a logical thermostat that only exists programmatically in blockchain. This could be connected to a real thermostat like Nest via API. You will create an asset, `Sensor`, that has a `teamID`, `teamName`, `sensorTemp`, `thermostatTemp` and `recommendation`. The `sensorTemp` value will be supplied by the Raspberry Pi Sens Hat Simulator through NodeRED to the `SetSensorTemp` transaction. To keep family members, housemates, friends or children from excessively running air conditioning or heat, they must first find out if they have permission to adjust the thermostat by running a transaction, `ChangeThermostatTemp`, defined in a smart contract running on Hyperledger Fabric. We will also add in the ability to consult current conditions via API from WeatherUnderground.com to optimally set the thermostat through the `CompareWeater` transaction.





In this lab you'll create the models for:

1 Asset: Sensor

3 Transactions:
* `SetSensorTemp` -- allows the Raspberry Pi to send the temperature from its sensor.
* `ChangeThermostatTemp` -- if the requested `thermostatTemp` value passed into the transaction is within 3 degrees of the `sensorTemp` then the `thermostatTemp` as part of the `Team` asset will be changed. Otherwise an error message is given.
* `CompareWeather` -- from the WeatherUnderground.com API, the `feelsLike` and `outsideTemp` are passed into the transaction. Depending on where the values fall, a `recommendation` for the option thermostat setting will be passed back to the `Team` asset and the `thermostatTemp` value for the `Team` asset will be updated.
