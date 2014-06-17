#!/usr/bin/env ruby
# Generated by the protocol buffer compiler. DO NOT EDIT!

require 'protocol_buffers'

module Main
  # forward declarations
  class GetProjectRootRequest < ::ProtocolBuffers::Message;
  end
  class GetProjectRootResponse < ::ProtocolBuffers::Message;
  end
  class GetAllStepsRequest < ::ProtocolBuffers::Message;
  end
  class GetAllStepsResponse < ::ProtocolBuffers::Message;
  end
  class GetAllSpecsRequest < ::ProtocolBuffers::Message;
  end
  class GetAllSpecsResponse < ::ProtocolBuffers::Message;
  end
  class GetStepValueRequest < ::ProtocolBuffers::Message;
  end
  class GetStepValueResponse < ::ProtocolBuffers::Message;
  end
  class ErrorResponse < ::ProtocolBuffers::Message;
  end
  class ProtoSpec < ::ProtocolBuffers::Message;
  end
  class ProtoItem < ::ProtocolBuffers::Message;
  end
  class ProtoHeading < ::ProtocolBuffers::Message;
  end
  class ProtoScenario < ::ProtocolBuffers::Message;
  end
  class ProtoStep < ::ProtocolBuffers::Message;
  end
  class ProtoConcept < ::ProtocolBuffers::Message;
  end
  class Fragment < ::ProtocolBuffers::Message;
  end
  class Parameter < ::ProtocolBuffers::Message;
  end
  class ProtoComment < ::ProtocolBuffers::Message;
  end
  class ProtoTableParam < ::ProtocolBuffers::Message;
  end
  class ProtoTableRow < ::ProtocolBuffers::Message;
  end
  class APIMessage < ::ProtocolBuffers::Message;
  end
  class ProtoTags < ::ProtocolBuffers::Message;
  end

  class GetProjectRootRequest < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetProjectRootRequest"

  end

  class GetProjectRootResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetProjectRootResponse"

    required :string, :projectRoot, 1
  end

  class GetAllStepsRequest < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetAllStepsRequest"

  end

  class GetAllStepsResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetAllStepsResponse"

    repeated :string, :steps, 1
  end

  class GetAllSpecsRequest < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetAllSpecsRequest"

  end

  class GetAllSpecsResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetAllSpecsResponse"

    repeated ::Main::ProtoSpec, :specs, 1
  end

  class GetStepValueRequest < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetStepValueRequest"

    required :string, :stepText, 1
  end

  class GetStepValueResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.GetStepValueResponse"

    required :string, :stepValue, 1
    repeated :string, :parameters, 2
  end

  class ErrorResponse < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.ErrorResponse"

    required :string, :error, 1
  end

  class ProtoSpec < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.ProtoSpec"

    repeated ::Main::ProtoItem, :items, 1
  end

  class ProtoItem < ::ProtocolBuffers::Message
    # forward declarations

    # enums
    module ItemType
      include ::ProtocolBuffers::Enum

      set_fully_qualified_name "main.ProtoItem.ItemType"

      Heading = 1
      Step = 2
      Concept = 4
      Scenario = 5
      Comment = 6
      Table = 7
      Tags = 8
    end

    set_fully_qualified_name "main.ProtoItem"

    required ::Main::ProtoItem::ItemType, :itemType, 1
    optional ::Main::ProtoHeading, :heading, 2
    optional ::Main::ProtoStep, :step, 3
    optional ::Main::ProtoConcept, :concept, 4
    optional ::Main::ProtoScenario, :scenario, 5
    optional ::Main::ProtoComment, :comment, 6
    optional ::Main::ProtoTableParam, :table, 7
    optional ::Main::ProtoTags, :tags, 8
  end

  class ProtoHeading < ::ProtocolBuffers::Message
    # forward declarations

    # enums
    module HeadingType
      include ::ProtocolBuffers::Enum

      set_fully_qualified_name "main.ProtoHeading.HeadingType"

      Spec = 1
      Scenario = 2
    end

    set_fully_qualified_name "main.ProtoHeading"

    required ::Main::ProtoHeading::HeadingType, :headingType, 1
    required :string, :text, 2
  end

  class ProtoScenario < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.ProtoScenario"

    repeated ::Main::ProtoItem, :scenarioItems, 1
  end

  class ProtoStep < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.ProtoStep"

    required :string, :text, 1
    repeated ::Main::Parameter, :parameters, 2
    repeated ::Main::Fragment, :fragments, 3
  end

  class ProtoConcept < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.ProtoConcept"

    required ::Main::ProtoStep, :conceptStep, 1
    repeated ::Main::ProtoStep, :steps, 2
  end

  class Fragment < ::ProtocolBuffers::Message
    # forward declarations

    # enums
    module FragmentType
      include ::ProtocolBuffers::Enum

      set_fully_qualified_name "main.Fragment.FragmentType"

      Text = 1
      Parameter = 2
    end

    set_fully_qualified_name "main.Fragment"

    required ::Main::Fragment::FragmentType, :fragmentType, 1
    optional :string, :text, 2
    optional ::Main::Parameter, :parameter, 3
  end

  class Parameter < ::ProtocolBuffers::Message
    # forward declarations

    # enums
    module ParameterType
      include ::ProtocolBuffers::Enum

      set_fully_qualified_name "main.Parameter.ParameterType"

      Static = 1
      Dynamic = 2
      Special = 3
      Table = 4
    end

    set_fully_qualified_name "main.Parameter"

    required ::Main::Parameter::ParameterType, :parameterType, 1
    optional :string, :value, 2
    optional ::Main::ProtoTableParam, :table, 3
  end

  class ProtoComment < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.ProtoComment"

    required :string, :text, 1
  end

  class ProtoTableParam < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.ProtoTableParam"

    required ::Main::ProtoTableRow, :headers, 1
    repeated ::Main::ProtoTableRow, :rows, 2
  end

  class ProtoTableRow < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.ProtoTableRow"

    repeated :string, :cells, 1
  end

  class APIMessage < ::ProtocolBuffers::Message
    # forward declarations

    # enums
    module APIMessageType
      include ::ProtocolBuffers::Enum

      set_fully_qualified_name "main.APIMessage.APIMessageType"

      GetProjectRootRequest = 1
      GetProjectRootResponse = 2
      GetAllStepsRequest = 3
      GetAllStepResponse = 4
      GetAllSpecsRequest = 5
      GetAllSpecsResponse = 6
      GetStepValueRequest = 7
      GetStepValueResponse = 8
      ErrorResponse = 9
    end

    set_fully_qualified_name "main.APIMessage"

    required ::Main::APIMessage::APIMessageType, :messageType, 1
    required :int64, :messageId, 2
    optional ::Main::GetProjectRootRequest, :projectRootRequest, 3
    optional ::Main::GetProjectRootResponse, :projectRootResponse, 4
    optional ::Main::GetAllStepsRequest, :allStepsRequest, 5
    optional ::Main::GetAllStepsResponse, :allStepsResponse, 6
    optional ::Main::GetAllSpecsRequest, :allSpecsRequest, 7
    optional ::Main::GetAllSpecsResponse, :allSpecsResponse, 8
    optional ::Main::GetStepValueRequest, :stepValueRequest, 9
    optional ::Main::GetStepValueResponse, :stepValueResponse, 10
    optional ::Main::ErrorResponse, :error, 11
  end

  class ProtoTags < ::ProtocolBuffers::Message
    set_fully_qualified_name "main.ProtoTags"

    repeated :string, :tags, 1
  end

end