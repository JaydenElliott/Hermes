# ArcStack Websocket Service

## Code logic

### Top level manager

- The entire application is served by the chatserver manager class.
- This class instantiates the channel manager and user manager
  <br>
  <br>

### Sub managers

- channel and user code is split into two sections: logic and models
- all interactions should go through the model files, which will interact with the logic files privately.
- this ensures variables are private, cannot be modified by outside sources and can be expanded upon without refactoring all the code entirely.
