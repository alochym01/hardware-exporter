# hardware-exporter
- hpilo_agentless_management_service_status
- hpilo_array_controller_status
- hpilo_avgcpu0freq
- hpilo_avgcpu1freq
- hpilo_bios_or_hardware_health_status
- hpilo_cpu0power
- hpilo_cpu1power
- hpilo_cpuicutil
- hpilo_cpuutil
- hpilo_fan_redundancy_status
- hpilo_fan_speed
- hpilo_fan_status
- hpilo_fans_status
- hpilo_ilo_port_status
- hpilo_iobusutil
- hpilo_jittercount
- hpilo_logical_drive_status
- hpilo_memories_status
- hpilo_memorybusutil
- hpilo_network_adapter_status
- hpilo_network_port_status
- hpilo_networks_status
- hpilo_physical_drive_metrics
- hpilo_physical_drive_power_on_hours
- hpilo_physical_drive_ssd_endurance
- hpilo_physical_drive_status
- hpilo_physical_drive_uncorrected_read_errors
- hpilo_physical_drive_uncorrected_write_errors
- hpilo_power_consumed_by_all
- hpilo_power_consumed_by_each
- hpilo_power_control
- hpilo_power_line_input_voltage
- hpilo_power_supplies_status
- hpilo_power_supply_redundancy_status
- hpilo_power_supply_status
- hpilo_processor_status
- hpilo_processors_status
- hpilo_smart_storage_batteries_status
- hpilo_smart_storage_battery_charge_level_percent
- hpilo_smart_storage_battery_status
- hpilo_storage_enclosure_status
- hpilo_storages_status
- hpilo_temperature_reading
- hpilo_temperature_reading_status
- hpilo_temperatures_status

- hpilo_memory_status - ok // Systems
  ```json
    // hpe
      "MemorySummary": {
        "Status": {
            "HealthRollup": "OK"
        },
        "TotalSystemMemoryGiB": 128,
        "TotalSystemPersistentMemoryGiB": 0
    }
    // dell
    "MemorySummary": {
        "MemoryMirroring": "System",
        "Status": {
            "Health": "OK",
            "HealthRollup": "OK",
            "State": "Enabled"
        },
        "TotalSystemMemoryGiB": 128
    },
  ```
- hpilo_system_health - ok // Systems
  ```json
  // hpe
  "Status": {
        "Health": "OK",
        "HealthRollup": "OK",
        "State": "Enabled"
    }
  // dell
  "Status": {
        "Health": "OK",
        "HealthRollup": "OK",
        "State": "Enabled"
    }
  ```
-
