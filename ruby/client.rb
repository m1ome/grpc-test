this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'service_services_pb'
require 'rack'

app = Proc.new do |env|
  stub = Greeter::Greeter::Stub.new('localhost:50000', :this_channel_is_insecure)
  counter = stub.get_counter(Greeter::Empty.new()).counter

  ['200', {'Content-Type' => 'text/html'}, ["Current count: #{counter}"]]
end
 
Rack::Handler::WEBrick.run app