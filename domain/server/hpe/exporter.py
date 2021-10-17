"""
Pulls data from specified iLO and presents as Prometheus metrics
"""
from __future__ import print_function
from _socket import gaierror
import sys
import json

import time
import prometheus_metrics
from BaseHTTPServer import BaseHTTPRequestHandler
from BaseHTTPServer import HTTPServer
from SocketServer import ForkingMixIn
from prometheus_client import generate_latest, Summary
from urlparse import parse_qs
from urlparse import urlparse

from _redfishobject import RedfishObject
from redfish.rest.v1 import ServerDownOrUnreachableError

def print_err(*args, **kwargs):
    print(*args, file=sys.stderr, **kwargs)

# Create a metric to track time spent and requests made.
REQUEST_TIME = Summary(
    'request_processing_seconds', 'Time spent processing request')


class ForkingHTTPServer(ForkingMixIn, HTTPServer):
    max_children = 30
    timeout = 30


class RequestHandler(BaseHTTPRequestHandler):
    """
    Endpoint handler
    """
    def return_error(self):
        self.send_response(500)
        self.end_headers()

    def do_GET(self):
        """
        Process GET request
        :return: Response with Prometheus metrics
        """
        # this will be used to return the total amount of time the request took
        start_time = time.time()
        # get parameters from the URL
        url = urlparse(self.path)
        # following boolean will be passed to True if an error is detected during the argument parsing
        error_detected = False
        query_components = parse_qs(urlparse(self.path).query)

        if self.path == '/favicon.ico':
            return

      #  ilo_host = "https://192.168.2.130"
        ilo_user = "username"
        ilo_password = "password"
        try:
            ilo_host = query_components['ilo_host'][0]
       #     ilo_user = query_components['ilo_user'][0]
        #    ilo_password = query_components['ilo_password'][0]
        except KeyError, e:
            print_err("missing parameter %s" % e)
            self.return_error()
            error_detected = True

        if url.path == self.server.endpoint and ilo_host and ilo_user and ilo_password:

            # Create a redfish object
            redfishobj = None
    	    try:
                redfishobj = RedfishObject(ilo_host, ilo_user, ilo_password)
    	    except ServerDownOrUnreachableError as excp:
        	sys.stderr.write("ERROR: server not reachable or doesn't support " \
                                                                "RedFish.\n")
        	sys.exit()
  	    except Exception as excp:
                raise excp

            # /redfish/v1/systems/1/
            systems = redfishobj.search_for_type("ComputerSystem.")
            system_rsp = {}
            for system in systems:
                system_rsp = redfishobj.redfish_get(system["@odata.id"])

            #print(system_rsp)
            #print('\n')

	    #system_rsp = system_rsp.dict

            BiosVersion = system_rsp.dict['BiosVersion']
            ServerModel = system_rsp.dict['Manufacturer'] + ' ' + system_rsp.dict['Model']
            PowerState = system_rsp.dict['PowerState']
            Bios_backup = system_rsp.dict['Oem']['Hpe']['Bios']['Backup']['VersionString']
            OsName = system_rsp.dict['Oem']['Hpe']['HostOS']['OsName'] + ' ' + system_rsp.dict['Oem']['Hpe']['HostOS']['OsVersion']
            OsSysDescription = system_rsp.dict['Oem']['Hpe']['HostOS']['OsSysDescription']
            IntelligentProvisioningVersion = system_rsp.dict['Oem']['Hpe']['IntelligentProvisioningVersion']
            PCAPartNumber = system_rsp.dict['Oem']['Hpe']['PCAPartNumber']
            PostState = system_rsp.dict['Oem']['Hpe']['PostState']
            SKU = system_rsp.dict['SKU']
            SerialNumber = system_rsp.dict['SerialNumber']
            SystemHealth = system_rsp.dict["Status"]["Health"]

            SystemHealth_temp = prometheus_metrics.hardware_system_health_gauge.labels(bios_version=BiosVersion,server_model=ServerModel,power_state=PowerState,bios_backup=Bios_backup,os_name=OsName,os_description=OsSysDescription,intelligent_provisioning_version=IntelligentProvisioningVersion,pca_part_number=PCAPartNumber,post_state=PostState,sku=SKU,serial_number=SerialNumber)

            if SystemHealth.upper() == 'OK':
                SystemHealth_temp.set(0)
            elif SystemHealth.upper() == 'WARNING':
                SystemHealth_temp.set(1)
            else:
                SystemHealth_temp.set(2)


            # AggregateHealthStatus - each component status
            AggregateHealthStatus = system_rsp.dict["Oem"]["Hpe"]["AggregateHealthStatus"]

            status_dict = {
                'Agentless_Management_Service': AggregateHealthStatus["AgentlessManagementService"],
                'Bios_Or_Hardware_Health': AggregateHealthStatus["BiosOrHardwareHealth"]["Status"]["Health"],
                'Fan_Redundancy': AggregateHealthStatus["FanRedundancy"],
                'Fans': AggregateHealthStatus["Fans"]["Status"]["Health"],
                'Memories': AggregateHealthStatus["Memory"]["Status"]["Health"],
                'Networks': AggregateHealthStatus["Network"]["Status"]["Health"],
                'Power_Supplies': AggregateHealthStatus["PowerSupplies"]["Status"]["Health"],
                'Power_Supply_Redundancy': AggregateHealthStatus["PowerSupplyRedundancy"],
                'Processors': AggregateHealthStatus["Processors"]["Status"]["Health"],
                'Smart_Storage_Batteries': AggregateHealthStatus["SmartStorageBattery"]["Status"]["Health"],
                'Storages': AggregateHealthStatus["Storage"]["Status"]["Health"],
                'Temperatures': AggregateHealthStatus["Temperatures"]["Status"]["Health"]
                }

            for key, value in status_dict.items():
                status_gauge = 'hpilo_{}_status_gauge'.format(key.lower())
                if value.upper() == 'OK' or value.upper() == 'READY' or value.upper() == 'REDUNDANT':
                    prometheus_metrics.status_gauges[status_gauge].set(0)
                elif value.upper() == 'WARNING':
                    prometheus_metrics.status_gauges[status_gauge].set(1)
                else:
                    prometheus_metrics.status_gauges[status_gauge].set(2)


            # System Usage metrics
            SystemUsage = system_rsp.dict["Oem"]["Hpe"]["SystemUsage"]

            for key, value in SystemUsage.items():
                system_usage_gauge = 'hpilo_{}_gauge'.format(key.lower())
                prometheus_metrics.system_usage_gauges[system_usage_gauge].set(value)




            # /redfish/v1/Systems/1/SmartStorage/ArrayControllers/x/
            arrays = redfishobj.search_for_type("ArrayController.")

            for array in arrays:
                array_rsp = redfishobj.redfish_get(array["@odata.id"])

                #print(array_rsp)
                #print("\n")

                # an array
                Id_array = array_rsp.dict['Id']
                Model_array = array_rsp.dict['Model']
                Location_array = array_rsp.dict['Location']
                LocationFormat_array = array_rsp.dict['LocationFormat']
                InternalPortCount = array_rsp.dict['InternalPortCount']
                BackupPowerSourceStatus = array_rsp.dict['BackupPowerSourceStatus']
                ControllerPartNumber = array_rsp.dict['ControllerPartNumber']
                ControllerBoard = array_rsp.dict['ControllerBoard']['Status']['Health']
                CurrentOperatingMode = array_rsp.dict['CurrentOperatingMode']
                ReadCachePercent = array_rsp.dict['ReadCachePercent']
                FirmwareVersion_array = array_rsp.dict['FirmwareVersion']['Current']['VersionString']
                Status_array = array_rsp.dict['Status']['Health']

                UnconfiguredDrives_rsp = redfishobj.redfish_get(array_rsp.dict['Links']['UnconfiguredDrives']['@odata.id'])
                UnconfiguredDrives = UnconfiguredDrives_rsp.dict['Members@odata.count']

                Status_array_temp = prometheus_metrics.hpilo_array_controller_status_gauge.labels(id=Id_array,model=Model_array,location=LocationFormat_array + ': ' + Location_array,unconfigured_drives=UnconfiguredDrives,internal_port_count=InternalPortCount,backup_power_source_status=BackupPowerSourceStatus,part_number=ControllerPartNumber,controller_board=ControllerBoard,current_operating_mode=CurrentOperatingMode,read_cache_percent=ReadCachePercent,firmware_version=FirmwareVersion_array)

                if Status_array.upper() == 'OK':
                    Status_array_temp.set(0)
                elif Status_array.upper() == 'WARNING':
                    Status_array_temp.set(1)
                else:
                    Status_array_temp.set(2)


                # Logical drives of an array
                logical_drives = redfishobj.redfish_get(array_rsp.dict['Links']['LogicalDrives']['@odata.id'])
                for logical_drive in logical_drives.dict['Members']:
                    logical_drive_rsp = redfishobj.redfish_get(logical_drive["@odata.id"])

                    #print(logical_drive_rsp)
                    #print('\n')

                    Id_logical_drive = logical_drive_rsp.dict['Id']
                    Capacity_logical_drive = round(logical_drive_rsp.dict['CapacityMiB']/1024.0, 2)
                    MediaType_logical_drive = logical_drive_rsp.dict['MediaType']
                    Status_logical_drive = logical_drive_rsp.dict['Status']['Health']
                    Raid_logical_drive = logical_drive_rsp.dict['Raid']
                    InterfaceType_logical_drive = logical_drive_rsp.dict['InterfaceType']
                    LogicalDriveType = logical_drive_rsp.dict['LogicalDriveType']
                    AccelerationMethod_logical_drive = logical_drive_rsp.dict['AccelerationMethod']

                    LogicalDriveStatusReasons = []
                    for LogicalDriveStatusReason in logical_drive_rsp.dict['LogicalDriveStatusReasons']:
                        LogicalDriveStatusReasons.append(LogicalDriveStatusReason.encode('utf-8'))

                    DataDrives = redfishobj.redfish_get(logical_drive_rsp.dict['Links']['DataDrives']['@odata.id'])

                    DataDrive_array = []
                    for DataDrive in DataDrives.dict['Members']:
                        DataDrive_rsp = redfishobj.redfish_get(DataDrive["@odata.id"])

                        DataDrive_array.append(DataDrive_rsp.dict['Id'].encode('utf-8'))

                    if Capacity_logical_drive > 1024:
                        Capacity_logical_drive = round(Capacity_logical_drive/1024, 2)
                        Capacity_logical_drive = str(Capacity_logical_drive) + ' TiB'
                    else:
                        Capacity_logical_drive = str(Capacity_logical_drive) + ' GiB'

                    Status_logical_drive_temp = prometheus_metrics.hpilo_logical_drive_status_gauge.labels(id_array=Id_array,id=Id_logical_drive,capacity=Capacity_logical_drive,media_type=MediaType_logical_drive,raid=Raid_logical_drive,physical_drive_id=DataDrive_array,logical_drive_status_reasons=LogicalDriveStatusReasons,interface_type=InterfaceType_logical_drive,logical_drive_type=LogicalDriveType,acceleration_method=AccelerationMethod_logical_drive)

                    if Status_logical_drive.upper() == 'OK':
                        Status_logical_drive_temp.set(0)
                    elif Status_logical_drive.upper() == 'WARNING':
                        Status_logical_drive_temp.set(1)
                    else:
                        Status_logical_drive_temp.set(2)

                # Physical drives
                physical_drives = redfishobj.redfish_get(array_rsp.dict['Links']['PhysicalDrives']['@odata.id'])

                for physical_drive in physical_drives.dict['Members']:
                    physical_drive_rsp = redfishobj.redfish_get(physical_drive["@odata.id"])

                    #print(physical_drive_rsp)
                    #print('\n')

                    Id_physical_drive = physical_drive_rsp.dict['Id']
                    Capacity_physical_drive = physical_drive_rsp.dict['CapacityGB']
                    Temperature_physical_drive = physical_drive_rsp.dict['CurrentTemperatureCelsius']
                    MaxTemp_physical_drive = physical_drive_rsp.dict['MaximumTemperatureCelsius']
                    Location_physical_drive = physical_drive_rsp.dict['LocationFormat'] + ': ' + physical_drive_rsp.dict['Location']
                    MediaType_physical_drive = physical_drive_rsp.dict['MediaType']
                    PowerOnHours_physical_drive = physical_drive_rsp.dict['PowerOnHours']
                    SSDEndurance = physical_drive_rsp.dict['SSDEnduranceUtilizationPercentage']
                    Model_physical_drive = physical_drive_rsp.dict['Model']
                    Status_physical_drive = physical_drive_rsp.dict['Status']['Health']
                    UncorrectedReadErrors = physical_drive_rsp.dict['UncorrectedReadErrors']
                    UncorrectedWriteErrors = physical_drive_rsp.dict['UncorrectedWriteErrors']
                    InterfaceType_physical_drive = physical_drive_rsp.dict['InterfaceType']
                    CarrierAuthenticationStatus = physical_drive_rsp.dict['CarrierAuthenticationStatus']
                    DiskDriveUse = physical_drive_rsp.dict['DiskDriveUse']
                    InterfaceSpeedMbps = physical_drive_rsp.dict['InterfaceSpeedMbps']
                    FirmwareVersion_physical_drive = physical_drive_rsp.dict['FirmwareVersion']['Current']['VersionString']
                    try:
                        RotationalSpeedRpm = physical_drive_rsp.dict['RotationalSpeedRpm']
                    except:
                        RotationalSpeedRpm = ''

                    DiskDriveStatusReasons = []
                    for DiskDriveStatusReason in physical_drive_rsp.dict['DiskDriveStatusReasons']:
                        DiskDriveStatusReasons.append(DiskDriveStatusReason.encode('utf-8'))


                    if Capacity_physical_drive > 1000:
                        Capacity_physical_drive = round(Capacity_physical_drive/1000, 2)
                        Capacity_physical_drive = str(Capacity_physical_drive) + ' TB'
                    else:
                        Capacity_physical_drive = str(Capacity_physical_drive) + ' GB'

                    Status_physical_drive_temp = prometheus_metrics.hpilo_physical_drive_status_gauge.labels(id=Id_physical_drive,location=Location_physical_drive,capacity=Capacity_physical_drive,media_type=MediaType_physical_drive,model=Model_physical_drive,id_array=Id_array,interface_type=InterfaceType_physical_drive,carrier_authentication_status=CarrierAuthenticationStatus,disk_drive_use=DiskDriveUse,interface_speed_mbps=InterfaceSpeedMbps,disk_drive_status_reasons=DiskDriveStatusReasons,firmware_version=FirmwareVersion_physical_drive,rotational_speed_rpm=RotationalSpeedRpm)

                    if Status_physical_drive.upper() == 'OK':
                        Status_physical_drive_temp.set(0)
                    elif Status_physical_drive.upper() == 'WARNING':
                        Status_physical_drive_temp.set(1)
                    else:
                        Status_physical_drive_temp.set(2)

                    prometheus_metrics.hpilo_physical_drive_metrics_gauge.labels(id=Id_physical_drive,maximum_recommended=MaxTemp_physical_drive,power_on_hours=PowerOnHours_physical_drive,ssd_endurance=SSDEndurance,uncorrected_read_errors=UncorrectedReadErrors,uncorrected_write_errors=UncorrectedWriteErrors,current_temperature=Temperature_physical_drive).set(0)

                    if PowerOnHours_physical_drive != None:
                        prometheus_metrics.hpilo_physical_drive_power_on_hours_gauge.labels(id=Id_physical_drive,location=Location_physical_drive,capacity=Capacity_physical_drive,media_type=MediaType_physical_drive,model=Model_physical_drive,id_array=Id_array,interface_type=InterfaceType_physical_drive,interface_speed_mbps=InterfaceSpeedMbps).set(PowerOnHours_physical_drive)

                    if SSDEndurance != None:
                        prometheus_metrics.hpilo_physical_drive_ssd_endurance_gauge.labels(id=Id_physical_drive,location=Location_physical_drive,capacity=Capacity_physical_drive,media_type=MediaType_physical_drive,model=Model_physical_drive,id_array=Id_array,interface_type=InterfaceType_physical_drive,interface_speed_mbps=InterfaceSpeedMbps).set(SSDEndurance)

                    prometheus_metrics.hpilo_physical_drive_uncorrected_read_errors_gauge.labels(id=Id_physical_drive,location=Location_physical_drive,capacity=Capacity_physical_drive,media_type=MediaType_physical_drive,model=Model_physical_drive,id_array=Id_array,interface_type=InterfaceType_physical_drive,interface_speed_mbps=InterfaceSpeedMbps).set(UncorrectedReadErrors)

                    prometheus_metrics.hpilo_physical_drive_uncorrected_write_errors_gauge.labels(id=Id_physical_drive,location=Location_physical_drive,capacity=Capacity_physical_drive,media_type=MediaType_physical_drive,model=Model_physical_drive,id_array=Id_array,interface_type=InterfaceType_physical_drive,interface_speed_mbps=InterfaceSpeedMbps).set(UncorrectedWriteErrors)


                # StorageEnclosures
                StorageEnclosures = redfishobj.redfish_get(array_rsp.dict['Links']['StorageEnclosures']['@odata.id'])


                for StorageEnclosure in StorageEnclosures.dict['Members']:
                    StorageEnclosure_rsp = redfishobj.redfish_get(StorageEnclosure["@odata.id"])

                    #print(StorageEnclosure_rsp)
                    #print('\n')

                    Id_StorageEnclosure = StorageEnclosure_rsp.dict['Id']
                    DriveBayCount = StorageEnclosure_rsp.dict['DriveBayCount']
                    Location_StorageEnclosure = StorageEnclosure_rsp.dict['LocationFormat'] + ': ' + StorageEnclosure_rsp.dict['Location']
                    Model_StorageEnclosure = StorageEnclosure_rsp.dict['Model']
                    FirmwareVersion_StorageEnclosure = StorageEnclosure_rsp.dict['FirmwareVersion']['Current']['VersionString']
                    Status_StorageEnclosure = StorageEnclosure_rsp.dict['Status']['Health']

                    Status_StorageEnclosure_temp = prometheus_metrics.hpilo_storage_enclosure_status_gauge.labels(id=Id_StorageEnclosure,location_identifier=Location_StorageEnclosure,model=Model_StorageEnclosure,id_array=Id_array,drive_bays_supported=DriveBayCount,firmware_version=FirmwareVersion_StorageEnclosure)

                    if Status_StorageEnclosure.upper() == 'OK':
                        Status_StorageEnclosure_temp.set(0)
                    elif Status_StorageEnclosure.upper() == 'WARNING':
                        Status_StorageEnclosure_temp.set(1)
                    else:
                        Status_StorageEnclosure_temp.set(2)


            # Memory
            Memorys = redfishobj.redfish_get('/redfish/v1/Systems/1/Memory/')


            for Memory in Memorys.dict['Members']:
                Memory_rsp = redfishobj.redfish_get(Memory["@odata.id"])

                #print(Memory_rsp)
                #print('\n')

                Location_Memory = Memory_rsp.dict['DeviceLocator']
                Channel_Memory = Memory_rsp.dict['MemoryLocation']['Channel']
                MemoryModuleStatus = Memory_rsp.dict['Oem']['Hpe']['DIMMStatus']
                Status_Memory = Memory_rsp.dict['Status']['Health']
                if MemoryModuleStatus != 'NotPresent' and MemoryModuleStatus != 'Null' and MemoryModuleStatus != 'Unknown' and MemoryModuleStatus != 'Other':
                    PartNumber_Memory = Memory_rsp.dict['PartNumber']
                    Capacity_Memory = str(Memory_rsp.dict['CapacityMiB']/1024) + ' GB'
                    MemoryDeviceType = Memory_rsp.dict['MemoryDeviceType']
                    OperatingSpeedMhz_Memory = Memory_rsp.dict['OperatingSpeedMhz']
                else:
                    PartNumber_Memory = ''
                    Capacity_Memory = ''
                    MemoryDeviceType = ''
                    OperatingSpeedMhz_Memory = ''

                Status_Memory_temp = prometheus_metrics.hpilo_memory_status_gauge.labels(location=Location_Memory,channel=Channel_Memory,module_status=MemoryModuleStatus,part_number=PartNumber_Memory,capacity=Capacity_Memory,device_type=MemoryDeviceType,operating_speed_mhz=OperatingSpeedMhz_Memory)

                if Status_Memory.upper() == 'OK':
                    Status_Memory_temp.set(0)
                elif Status_Memory.upper() == 'WARNING':
                    Status_Memory_temp.set(1)
                else:
                    Status_Memory_temp.set(2)


            # Processors
            Processors = redfishobj.redfish_get('/redfish/v1/Systems/1/Processors/')

            for Processor in Processors.dict['Members']:
                Processor_rsp = redfishobj.redfish_get(Processor["@odata.id"])

                #print(Processor_rsp)
                #print('\n')

                Id_cpu = Processor_rsp.dict['Id']
                Model_cpu = Processor_rsp.dict['Model']

                Caches_cpu = []
                for Cache_cpu_rsp in Processor_rsp.dict['Oem']['Hpe']['Cache']:
                    Cache_cpu = Cache_cpu_rsp['Name'] + ': ' + str(round(Cache_cpu_rsp['InstalledSizeKB']/1024.0, 2)) + ' MB'
                    Caches_cpu.append(Cache_cpu.encode('utf-8'))

                Cores_Threads = str(Processor_rsp.dict['TotalCores']) + '/' + str(Processor_rsp.dict['TotalThreads'])
                LastBootSpeed = str(Processor_rsp.dict['Oem']['Hpe']['RatedSpeedMHz']) + ' MHz'
                MaxSpeedMHz = str(Processor_rsp.dict['MaxSpeedMHz']) + ' MHz'
                CoresEnabled = Processor_rsp.dict['Oem']['Hpe']['CoresEnabled']
                Status_cpu = Processor_rsp.dict['Status']['Health']

                Status_cpu_temp = prometheus_metrics.hpilo_processor_status_gauge.labels(id=Id_cpu,model=Model_cpu,caches=Caches_cpu,cores_threads=Cores_Threads,last_boot_speed=LastBootSpeed,max_speed_supported=MaxSpeedMHz,cores_enabled=CoresEnabled)

                if Status_cpu.upper() == 'OK':
                    Status_cpu_temp.set(0)
                elif Status_cpu.upper() == 'WARNING':
                    Status_cpu_temp.set(1)
                else:
                    Status_cpu_temp.set(2)

            # Chassis

            chassis_rsp = redfishobj.redfish_get('/redfish/v1/Chassis/1/')

            #print(chassis_rsp)
            #print('\n')

            for SmartStorageBattery in chassis_rsp.dict['Oem']['Hpe']['SmartStorageBattery']:
                Id_battery = SmartStorageBattery['Index']
                ChargeLevelPercent = SmartStorageBattery['ChargeLevelPercent']
                MaximumCapWatts = SmartStorageBattery['MaximumCapWatts']
                Model_battery = SmartStorageBattery['Model']
                RemainingChargeTimeSeconds = SmartStorageBattery['RemainingChargeTimeSeconds']
                SparePartNumber = SmartStorageBattery['SparePartNumber']
                FirmwareVersion_battery = SmartStorageBattery['FirmwareVersion']

                Status_battery = SmartStorageBattery['Status']['Health']

                prometheus_metrics.hpilo_smart_storage_battery_charge_level_percent_gauge.labels(id=Id_battery,model=Model_battery,spare_part_number=SparePartNumber,firmware_version=FirmwareVersion_battery).set(ChargeLevelPercent)

                Status_battery_temp = prometheus_metrics.hpilo_smart_storage_battery_status_gauge.labels(id=Id_battery,charge_level_percent=ChargeLevelPercent,maximum_capacity_watts=MaximumCapWatts,model=Model_battery,remaining_charge_time_seconds=RemainingChargeTimeSeconds,spare_part_number=SparePartNumber,firmware_version=FirmwareVersion_battery)

                if Status_battery.upper() == 'OK':
                    Status_battery_temp.set(0)
                elif Status_battery.upper() == 'WARNING':
                    Status_battery_temp.set(1)
                else:
                    Status_battery_temp.set(2)



            # powers in chassis
            powers_rsp = redfishobj.redfish_get(chassis_rsp.dict['Power']['@odata.id'])

            #print(powers_rsp)
            #print('\n')

            # power control
            for power_control in powers_rsp.dict['PowerControl']:

                Id_power_control = power_control['MemberId']
                PowerCapacityWatts = power_control['PowerCapacityWatts']
                PowerConsumedWatts = power_control['PowerConsumedWatts']
                MaxConsumedWatts = power_control['PowerMetrics']['MaxConsumedWatts']
                MinConsumedWatts = power_control['PowerMetrics']['MinConsumedWatts']
                AverageConsumedWatts = power_control['PowerMetrics']['AverageConsumedWatts']
                IntervalInMin = power_control['PowerMetrics']['IntervalInMin']

                prometheus_metrics.hpilo_power_consumed_by_all_gauge.labels(id=Id_power_control,capacity=PowerCapacityWatts).set(PowerConsumedWatts)

                prometheus_metrics.hpilo_power_control_gauge.labels(id=Id_power_control,capacity=PowerCapacityWatts,max_consumed=MaxConsumedWatts,min_consumed=MinConsumedWatts,average_consumed=AverageConsumedWatts,interval_in_min=IntervalInMin).set(0)


            # power supply
            for power_supply in powers_rsp.dict['PowerSupplies']:

                Id_power_supply = power_supply['MemberId']
                Model_power_supply = power_supply['Model']
                SparePartNumber_power_supply = power_supply['SparePartNumber']
                FirmwareVersion_power_supply = power_supply['FirmwareVersion']
                Location_power_supply = 'BayNumber: ' + str(power_supply['Oem']['Hpe']['BayNumber'])
                HotplugCapable = power_supply['Oem']['Hpe']['HotplugCapable']
                Capacity_power_supply = str(power_supply['PowerCapacityWatts']) + 'W'
                PowerSupplyStatus = power_supply['Oem']['Hpe']['PowerSupplyStatus']['State']
                MaxPowerOutputWatts = power_supply['Oem']['Hpe']['MaxPowerOutputWatts']

                Status_power_supply = power_supply['Status']['Health']
                LineInputVoltage = power_supply['LineInputVoltage']
                AveragePowerOutputWatts = power_supply['Oem']['Hpe']['AveragePowerOutputWatts']

                prometheus_metrics.hpilo_power_consumed_by_each_gauge.labels(id=Id_power_supply,capacity=Capacity_power_supply,model=Model_power_supply,location=Location_power_supply).set(AveragePowerOutputWatts)

                prometheus_metrics.hpilo_power_line_input_voltage_gauge.labels(id=Id_power_supply,capacity=Capacity_power_supply,model=Model_power_supply,location=Location_power_supply).set(LineInputVoltage)

                Status_power_supply_temp = prometheus_metrics.hpilo_power_supply_status_gauge.labels(id=Id_power_supply,capacity=Capacity_power_supply,model=Model_power_supply,location=Location_power_supply,hot_plug_capable=HotplugCapable,power_supply_status=PowerSupplyStatus,max_output_watts_10s_interval=MaxPowerOutputWatts,spare_part_number=SparePartNumber_power_supply,firmware_version=FirmwareVersion_power_supply)

                if Status_power_supply.upper() == 'OK':
                    Status_power_supply_temp.set(0)
                elif Status_power_supply.upper() == 'WARNING':
                    Status_power_supply_temp.set(1)
                else:
                    Status_power_supply_temp.set(2)

            # thermals in chassis
            thermals_rsp = redfishobj.redfish_get(chassis_rsp.dict['Thermal']['@odata.id'])

            #print(thermals_rsp)
            #print('\n')

            # fans in thermals
            for fan in thermals_rsp.dict['Fans']:

                Id_fan = fan['MemberId']
                Name_fan = fan['Name']
                HotPluggable_fan = fan['Oem']['Hpe']['HotPluggable']
                Location_fan = fan['Oem']['Hpe']['Location']

                Reading_fan = fan['Reading']
                Status_fan = fan['Status']['Health']

                prometheus_metrics.hpilo_fan_speed_gauge.labels(id=Id_fan,name=Name_fan,hot_pluggable=HotPluggable_fan,cool_location=Location_fan).set(Reading_fan)

                Status_fan_temp = prometheus_metrics.hpilo_fan_status_gauge.labels(id=Id_fan,name=Name_fan,hot_pluggable=HotPluggable_fan,cool_location=Location_fan)

                if Status_fan.upper() == 'OK':
                    Status_fan_temp.set(0)
                elif Status_fan.upper() == 'WARNING':
                    Status_fan_temp.set(1)
                else:
                    Status_fan_temp.set(2)

            # temperatures in thermals
            for temperature in thermals_rsp.dict['Temperatures']:

                if temperature['Status']['State'] == 'Enabled':
                    Id_temperature = temperature['MemberId']
                    Name_temperature = temperature['Name']
                    PhysicalContext = temperature['PhysicalContext']
                    UpperThresholdCritical = temperature['UpperThresholdCritical']
                    UpperThresholdFatal = temperature['UpperThresholdFatal']

                    Reading_temperature = temperature['ReadingCelsius']
                    Status_temperature = temperature['Status']['Health']

                    prometheus_metrics.hpilo_temperature_reading_gauge.labels(id=Id_temperature,name=Name_temperature,location=PhysicalContext,caution=UpperThresholdCritical,critical=UpperThresholdFatal).set(Reading_temperature)

                    Status_temperature_temp = prometheus_metrics.hpilo_temperature_reading_status_gauge.labels(id=Id_temperature,name=Name_temperature,location=PhysicalContext,caution=UpperThresholdCritical,critical=UpperThresholdFatal,temperature=Reading_temperature)

                    if Status_temperature.upper() == 'OK':
                        Status_temperature_temp.set(0)
                    elif Status_temperature.upper() == 'WARNING':
                        Status_temperature_temp.set(1)
                    else:
                        Status_temperature_temp.set(2)



            # Network adapter
            BaseNetworkAdapters_rsp = redfishobj.redfish_get('/redfish/v1/Systems/1/BaseNetworkAdapters/')

            for BaseNetworkAdapter in BaseNetworkAdapters_rsp.dict['Members']:
                BaseNetworkAdapter_rsp = redfishobj.redfish_get(BaseNetworkAdapter['@odata.id'])

                #print(BaseNetworkAdapter_rsp)
                #print('\n')

                Id_network_adapter = BaseNetworkAdapter_rsp.dict['Id']
                Name_network_adapter = BaseNetworkAdapter_rsp.dict['Name']
                FirmwareVersion_network_adapter = BaseNetworkAdapter_rsp.dict['Firmware']['Current']['VersionString']
                try:
                    PartNumber_network_adapter = BaseNetworkAdapter_rsp.dict['PartNumber']
                except:
                    PartNumber_network_adapter = ''

                try:
                    Status_network_adapter = BaseNetworkAdapter_rsp.dict['Status']['Health']
                except:
                    Status_network_adapter = ''

                Status_network_adapter_temp = prometheus_metrics.hpilo_network_adapter_status_gauge.labels(id=Id_network_adapter,name=Name_network_adapter,part_number=PartNumber_network_adapter,firmware_version=FirmwareVersion_network_adapter)

                if Status_network_adapter.upper() == 'OK':
                    Status_network_adapter_temp.set(0)
                elif Status_network_adapter.upper() == 'WARNING':
                    Status_network_adapter_temp.set(1)
                else:
                    Status_network_adapter_temp.set(2)


                # Network port
                for PhysicalPort in BaseNetworkAdapter_rsp.dict['PhysicalPorts']:

                    #print(PhysicalPort)
                    #print('\n')

                    Port_network = BaseNetworkAdapter_rsp.dict['PhysicalPorts'].index(PhysicalPort) + 1
                    FullDuplex = PhysicalPort['FullDuplex']

                    IPv4Addresses = []
                    for IPv4Address in PhysicalPort['IPv4Addresses']:
                        IPv4Addresses.append(IPv4Address['Address'].encode('utf-8'))

                    LinkStatus = PhysicalPort['LinkStatus']
                    MacAddress_network_port = PhysicalPort['MacAddress']
                    Name_network_port = PhysicalPort['Name']
                    SpeedMbps = PhysicalPort['SpeedMbps']

                    try:
                        Team_network_port = PhysicalPort['Oem']['Hpe']['Team']
                    except:
                        Team_network_port = ''

                    BadReceives = PhysicalPort['Oem']['Hpe']['BadReceives']
                    BadTransmits = PhysicalPort['Oem']['Hpe']['BadTransmits']
                    GoodReceives = PhysicalPort['Oem']['Hpe']['GoodReceives']
                    GoodTransmits = PhysicalPort['Oem']['Hpe']['GoodTransmits']

                    try:
                        Status_network_port = PhysicalPort['Status']['Health']
                    except:
                        Status_network_port=''

                    prometheus_metrics.hpilo_network_port_bad_receives_gauge.labels(id_adapter=Id_network_adapter,name_adapter=Name_network_adapter,port_number=Port_network,full_duplex=FullDuplex,ipv4_addresses=IPv4Addresses,link_status=LinkStatus,name=Name_network_port,speed_mbps=SpeedMbps,team=Team_network_port,mac_address=MacAddress_network_port).set(BadReceives)

                    prometheus_metrics.hpilo_network_port_bad_transmits_gauge.labels(id_adapter=Id_network_adapter,name_adapter=Name_network_adapter,port_number=Port_network,full_duplex=FullDuplex,ipv4_addresses=IPv4Addresses,link_status=LinkStatus,name=Name_network_port,speed_mbps=SpeedMbps,team=Team_network_port,mac_address=MacAddress_network_port).set(BadTransmits)

                    prometheus_metrics.hpilo_network_port_good_receives_gauge.labels(id_adapter=Id_network_adapter,name_adapter=Name_network_adapter,port_number=Port_network,full_duplex=FullDuplex,ipv4_addresses=IPv4Addresses,link_status=LinkStatus,name=Name_network_port,speed_mbps=SpeedMbps,team=Team_network_port,mac_address=MacAddress_network_port).set(GoodReceives)

                    prometheus_metrics.hpilo_network_port_good_transmits_gauge.labels(id_adapter=Id_network_adapter,name_adapter=Name_network_adapter,port_number=Port_network,full_duplex=FullDuplex,ipv4_addresses=IPv4Addresses,link_status=LinkStatus,name=Name_network_port,speed_mbps=SpeedMbps,team=Team_network_port,mac_address=MacAddress_network_port).set(GoodTransmits)

                    Status_network_port_temp = prometheus_metrics.hpilo_network_port_status_gauge.labels(id_adapter=Id_network_adapter,name_adapter=Name_network_adapter,port_number=Port_network,full_duplex=FullDuplex,ipv4_addresses=IPv4Addresses,link_status=LinkStatus,name=Name_network_port,speed_mbps=SpeedMbps,team=Team_network_port,mac_address=MacAddress_network_port)

                    if Status_network_port.upper() == 'OK':
                        Status_network_port_temp.set(0)
                    elif Status_network_port.upper() == 'WARNING':
                        Status_network_port_temp.set(1)
                    else:
                        Status_network_port_temp.set(2)

            # ilo port
            ilo_port_rsp = redfishobj.redfish_get('/redfish/v1/Managers/1/EthernetInterfaces/1/')

            #print(ilo_port_rsp)
            #print('\n')

            FullDuplex_ilo = ilo_port_rsp.dict['FullDuplex']
            Address_ilo = ilo_port_rsp.dict['IPv4Addresses'][0]['Address']
            AddressOrigin_ilo = ilo_port_rsp.dict['IPv4Addresses'][0]['AddressOrigin']
            Gateway_ilo = ilo_port_rsp.dict['IPv4Addresses'][0]['Gateway']
            SubnetMask_ilo = ilo_port_rsp.dict['IPv4Addresses'][0]['SubnetMask']
            InterfaceEnabled_ilo = ilo_port_rsp.dict['InterfaceEnabled']
            InterfaceType_ilo = ilo_port_rsp.dict['Oem']['Hpe']['InterfaceType']
            SpeedMbps_ilo = ilo_port_rsp.dict['SpeedMbps']

            Status_ilo = ilo_port_rsp.dict['Status']['Health']

            Status_ilo_temp = prometheus_metrics.hpilo_ilo_port_status_gauge.labels(full_duplex=FullDuplex_ilo,address=Address_ilo,address_origin=AddressOrigin_ilo,gateway=Gateway_ilo,subnet_mask=SubnetMask_ilo,interface_enabled=InterfaceEnabled_ilo,interface_type=InterfaceType_ilo,speed_mbps=SpeedMbps_ilo)

            if Status_ilo.upper() == 'OK':
                Status_ilo_temp.set(0)
            elif Status_ilo.upper() == 'WARNING':
                Status_ilo_temp.set(1)
            else:
                Status_ilo_temp.set(2)


            # get the amount of time the request took
            REQUEST_TIME.observe(time.time() - start_time)

            # generate and publish metrics
            metrics = generate_latest(prometheus_metrics.registry)
            self.send_response(200)
            self.send_header('Content-Type', 'text/plain')
            self.end_headers()
            self.wfile.write(metrics)

	else:
            if not error_detected:
                self.send_response(404)
                self.end_headers()


class ILOExporterServer(object):
    """
    Basic server implementation that exposes metrics to Prometheus
    """

    def __init__(self, address='0.0.0.0', port=8080, endpoint="/metrics"):
        self._address = address
        self._port = port
        self.endpoint = endpoint

    def print_info(self):
        print_err("Starting exporter on: http://{}:{}{}".format(self._address, self._port, self.endpoint))
        print_err("Press Ctrl+C to quit")

    def run(self):
        self.print_info()

        server = ForkingHTTPServer((self._address, self._port), RequestHandler)
        server.endpoint = self.endpoint

        try:
            while True:
                server.handle_request()
        except KeyboardInterrupt:
            print_err("Killing exporter")
            server.server_close()
