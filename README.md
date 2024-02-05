## Effortless Hot Reloading in Golang: Harnessing the Power of Viper
#### fsnotify, and Callback Functions for Dynamic Configuration Updates
* <b>viper.OnConfigChange</b> method allows to register a callback function that gets executed when a change in the monitored configuration file is detected<br>
  Internally, Viper uses the fsnotify library to monitor the specified configuration file for changes