package hpe

import "github.com/alochym01/hardware-exporter/domain/server/base"

type PowerControlOem struct {
}
type ChassisPowerControlChassisPowerControl struct {
	AverageConsumedWatts int `json:"AverageConsumedWatts"`
	IntervalInMin        int `json:"IntervalInMin"`
	MaxConsumedWatts     int `json:"MaxConsumedWatts"`
	MinConsumedWatts     int `json:"MinConsumedWatts"`
}
type ChassisPowerControl struct {
	OdataID            string  `json:"@odata.id"`
	MemberId           string  `json:"MemberId"`
	PowerCapacityWatts int     `json:"PowerCapacityWatts"`
	PowerConsumedWatts float64 `json:"PowerConsumedWatts"`
	PowerMetrics       ChassisPowerControlChassisPowerControl
	/*
		{
			"@odata.id": "/redfish/v1/Chassis/1/Power#PowerControl/0",
			"MemberId": "0",
			"PowerCapacityWatts": 1600,
			"PowerConsumedWatts": 412,
			"PowerMetrics": {
				"AverageConsumedWatts": 408,
				"IntervalInMin": 20,
				"MaxConsumedWatts": 478,
				"MinConsumedWatts": 404
			}
		}
	*/
}
type ChassisPowerSupplies struct {
	OdataID  string `json:"@odata.id"`
	MemberId string `json:"MemberId"`

	/*
		{
			"@odata.id": "/redfish/v1/Chassis/1/Power#PowerSupplies/0",
			"FirmwareVersion": "1.02",
			"LastPowerOutputWatts": 192,
			"LineInputVoltage": 228,
			"LineInputVoltageType": "ACHighLine",
			"Manufacturer": "CHCNY",
			"MemberId": "0",
			"Model": "865414-B21",
			"Name": "HpeServerPowerSupply",
			"Oem": {
				"Hpe": {
					"@odata.context": "/redfish/v1/$metadata#HpeServerPowerSupply.HpeServerPowerSupply",
					"@odata.type": "#HpeServerPowerSupply.v2_0_0.HpeServerPowerSupply",
					"AveragePowerOutputWatts": 192,
					"BayNumber": 1,
					"HotplugCapable": true,
					"MaxPowerOutputWatts": 219,
					"Mismatched": false,
					"PowerSupplyStatus": {
						"State": "Ok"
					},
					"iPDUCapable": false
				}
			}
		}
	*/
}
type ChassisRedundancy struct {
}
type PowerControl struct {
	base.Meta
	Id            string `json:"Id"`
	Name          string `json:"Name"`
	Oem           PowerControlOem
	PowerControl  []ChassisPowerControl
	PowerSupplies []ChassisPowerSupplies
	Redundancy    []ChassisRedundancy
	/*
	   {
	       "@odata.context": "/redfish/v1/$metadata#Power.Power",
	       "@odata.etag": "W/\"3F8B9A56\"",
	       "@odata.id": "/redfish/v1/Chassis/1/Power",
	       "@odata.type": "#Power.v1_3_0.Power",
	       "Id": "Power",
	       "Name": "PowerMetrics",
	       "Oem": {
	           "Hpe": {
	               "@odata.context": "/redfish/v1/$metadata#HpePowerMetricsExt.HpePowerMetricsExt",
	               "@odata.type": "#HpePowerMetricsExt.v2_2_0.HpePowerMetricsExt",
	               "BrownoutRecoveryEnabled": true,
	               "HasCpuPowerMetering": true,
	               "HasDimmPowerMetering": true,
	               "HasGpuPowerMetering": false,
	               "HasPowerMetering": true,
	               "HighEfficiencyMode": "Balanced",
	               "Links": {
	                   "FastPowerMeter": {
	                       "@odata.id": "/redfish/v1/Chassis/1/Power/FastPowerMeter"
	                   },
	                   "FederatedGroupCapping": {
	                       "@odata.id": "/redfish/v1/Chassis/1/Power/FederatedGroupCapping"
	                   },
	                   "PowerMeter": {
	                       "@odata.id": "/redfish/v1/Chassis/1/Power/PowerMeter"
	                   }
	               },
	               "MinimumSafelyAchievableCap": null,
	               "MinimumSafelyAchievableCapValid": false,
	               "SNMPPowerThresholdAlert": {
	                   "DurationInMin": 0,
	                   "ThresholdWatts": 0,
	                   "Trigger": "Disabled"
	               }
	           }
	       },
	       "PowerControl": [
	           {
	               "@odata.id": "/redfish/v1/Chassis/1/Power#PowerControl/0",
	               "MemberId": "0",
	               "PowerCapacityWatts": 1600,
	               "PowerConsumedWatts": 412,
	               "PowerMetrics": {
	                   "AverageConsumedWatts": 408,
	                   "IntervalInMin": 20,
	                   "MaxConsumedWatts": 478,
	                   "MinConsumedWatts": 404
	               }
	           }
	       ],
	       "PowerSupplies": [
	           {
	               "@odata.id": "/redfish/v1/Chassis/1/Power#PowerSupplies/0",
	               "FirmwareVersion": "1.02",
	               "LastPowerOutputWatts": 192,
	               "LineInputVoltage": 228,
	               "LineInputVoltageType": "ACHighLine",
	               "Manufacturer": "CHCNY",
	               "MemberId": "0",
	               "Model": "865414-B21",
	               "Name": "HpeServerPowerSupply",
	               "Oem": {
	                   "Hpe": {
	                       "@odata.context": "/redfish/v1/$metadata#HpeServerPowerSupply.HpeServerPowerSupply",
	                       "@odata.type": "#HpeServerPowerSupply.v2_0_0.HpeServerPowerSupply",
	                       "AveragePowerOutputWatts": 192,
	                       "BayNumber": 1,
	                       "HotplugCapable": true,
	                       "MaxPowerOutputWatts": 219,
	                       "Mismatched": false,
	                       "PowerSupplyStatus": {
	                           "State": "Ok"
	                       },
	                       "iPDUCapable": false
	                   }
	               },
	               "PowerCapacityWatts": 800,
	               "PowerSupplyType": "AC",
	               "SerialNumber": "5WEBP0D8JC665O",
	               "SparePartNumber": "866730-001",
	               "Status": {
	                   "Health": "OK",
	                   "State": "Enabled"
	               }
	           },
	           {
	               "@odata.id": "/redfish/v1/Chassis/1/Power#PowerSupplies/1",
	               "FirmwareVersion": "1.02",
	               "LastPowerOutputWatts": 220,
	               "LineInputVoltage": 229,
	               "LineInputVoltageType": "ACHighLine",
	               "Manufacturer": "CHCNY",
	               "MemberId": "1",
	               "Model": "865414-B21",
	               "Name": "HpeServerPowerSupply",
	               "Oem": {
	                   "Hpe": {
	                       "@odata.context": "/redfish/v1/$metadata#HpeServerPowerSupply.HpeServerPowerSupply",
	                       "@odata.type": "#HpeServerPowerSupply.v2_0_0.HpeServerPowerSupply",
	                       "AveragePowerOutputWatts": 220,
	                       "BayNumber": 2,
	                       "HotplugCapable": true,
	                       "MaxPowerOutputWatts": 243,
	                       "Mismatched": false,
	                       "PowerSupplyStatus": {
	                           "State": "Ok"
	                       },
	                       "iPDUCapable": false
	                   }
	               },
	               "PowerCapacityWatts": 800,
	               "PowerSupplyType": "AC",
	               "SerialNumber": "5WEBP0D8JC66BB",
	               "SparePartNumber": "866730-001",
	               "Status": {
	                   "Health": "OK",
	                   "State": "Enabled"
	               }
	           }
	       ],
	       "Redundancy": [
	           {
	               "@odata.id": "/redfish/v1/Chassis/1/Power#Redundancy/0",
	               "MaxNumSupported": 2,
	               "MemberId": "0",
	               "MinNumNeeded": 2,
	               "Mode": "Failover",
	               "Name": "PowerSupply Redundancy Group 1",
	               "RedundancySet": [
	                   {
	                       "@odata.id": "/redfish/v1/Chassis/1/Power#PowerSupplies/0"
	                   },
	                   {
	                       "@odata.id": "/redfish/v1/Chassis/1/Power#PowerSupplies/1"
	                   }
	               ],
	               "Status": {
	                   "Health": "OK",
	                   "State": "Enabled"
	               }
	           }
	       ]
	   }

	*/
}
