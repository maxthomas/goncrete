// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package goncrete

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type AnnotateCommunicationService interface {  //Annotator service methods. For concrete analytics that
  //are to be stood up as independent services, accessible
  //from any programming language.

  // Main annotation method. Takes a communication as input
  // and returns a new one as output.
  // 
  // It is up to the implementing service to verify that
  // the input communication is valid.
  // 
  // Can throw a ConcreteThriftException upon error
  // (invalid input, analytic exception, etc.).
  // 
  // Parameters:
  //  - Original
  Annotate(original *Communication) (r *Communication, err error)
  // Return the tool's AnnotationMetadata.
  GetMetadata() (r *AnnotationMetadata, err error)
  // Return a detailed description of what the particular tool
  // does, what inputs and outputs to expect, etc.
  // 
  // Developers whom are not familiar with the particular
  // analytic should be able to read this string and
  // understand the essential functions of the analytic.
  GetDocumentation() (r string, err error)
  // Indicate to the server it should shut down.
  Shutdown() (err error)
}

//Annotator service methods. For concrete analytics that
//are to be stood up as independent services, accessible
//from any programming language.
type AnnotateCommunicationServiceClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewAnnotateCommunicationServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AnnotateCommunicationServiceClient {
  return &AnnotateCommunicationServiceClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewAnnotateCommunicationServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AnnotateCommunicationServiceClient {
  return &AnnotateCommunicationServiceClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Main annotation method. Takes a communication as input
// and returns a new one as output.
// 
// It is up to the implementing service to verify that
// the input communication is valid.
// 
// Can throw a ConcreteThriftException upon error
// (invalid input, analytic exception, etc.).
// 
// Parameters:
//  - Original
func (p *AnnotateCommunicationServiceClient) Annotate(original *Communication) (r *Communication, err error) {
  if err = p.sendAnnotate(original); err != nil { return }
  return p.recvAnnotate()
}

func (p *AnnotateCommunicationServiceClient) sendAnnotate(original *Communication)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("annotate", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := AnnotateCommunicationServiceAnnotateArgs{
  Original : original,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *AnnotateCommunicationServiceClient) recvAnnotate() (value *Communication, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "annotate" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "annotate failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "annotate failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error1 error
    error1, err = error0.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error1
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "annotate failed: invalid message type")
    return
  }
  result := AnnotateCommunicationServiceAnnotateResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  if result.Ex != nil {
    err = result.Ex
    return 
  }
  value = result.GetSuccess()
  return
}

// Return the tool's AnnotationMetadata.
func (p *AnnotateCommunicationServiceClient) GetMetadata() (r *AnnotationMetadata, err error) {
  if err = p.sendGetMetadata(); err != nil { return }
  return p.recvGetMetadata()
}

func (p *AnnotateCommunicationServiceClient) sendGetMetadata()(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("getMetadata", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := AnnotateCommunicationServiceGetMetadataArgs{
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *AnnotateCommunicationServiceClient) recvGetMetadata() (value *AnnotationMetadata, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "getMetadata" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getMetadata failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getMetadata failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error3 error
    error3, err = error2.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error3
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getMetadata failed: invalid message type")
    return
  }
  result := AnnotateCommunicationServiceGetMetadataResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Return a detailed description of what the particular tool
// does, what inputs and outputs to expect, etc.
// 
// Developers whom are not familiar with the particular
// analytic should be able to read this string and
// understand the essential functions of the analytic.
func (p *AnnotateCommunicationServiceClient) GetDocumentation() (r string, err error) {
  if err = p.sendGetDocumentation(); err != nil { return }
  return p.recvGetDocumentation()
}

func (p *AnnotateCommunicationServiceClient) sendGetDocumentation()(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("getDocumentation", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := AnnotateCommunicationServiceGetDocumentationArgs{
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *AnnotateCommunicationServiceClient) recvGetDocumentation() (value string, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "getDocumentation" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getDocumentation failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getDocumentation failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error5 error
    error5, err = error4.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error5
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getDocumentation failed: invalid message type")
    return
  }
  result := AnnotateCommunicationServiceGetDocumentationResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Indicate to the server it should shut down.
func (p *AnnotateCommunicationServiceClient) Shutdown() (err error) {
  if err = p.sendShutdown(); err != nil { return }
  return
}

func (p *AnnotateCommunicationServiceClient) sendShutdown()(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("shutdown", thrift.ONEWAY, p.SeqId); err != nil {
      return
  }
  args := AnnotateCommunicationServiceShutdownArgs{
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


type AnnotateCommunicationServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler AnnotateCommunicationService
}

func (p *AnnotateCommunicationServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *AnnotateCommunicationServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *AnnotateCommunicationServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewAnnotateCommunicationServiceProcessor(handler AnnotateCommunicationService) *AnnotateCommunicationServiceProcessor {

  self6 := &AnnotateCommunicationServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self6.processorMap["annotate"] = &annotateCommunicationServiceProcessorAnnotate{handler:handler}
  self6.processorMap["getMetadata"] = &annotateCommunicationServiceProcessorGetMetadata{handler:handler}
  self6.processorMap["getDocumentation"] = &annotateCommunicationServiceProcessorGetDocumentation{handler:handler}
  self6.processorMap["shutdown"] = &annotateCommunicationServiceProcessorShutdown{handler:handler}
return self6
}

func (p *AnnotateCommunicationServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x7.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x7

}

type annotateCommunicationServiceProcessorAnnotate struct {
  handler AnnotateCommunicationService
}

func (p *annotateCommunicationServiceProcessorAnnotate) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := AnnotateCommunicationServiceAnnotateArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("annotate", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := AnnotateCommunicationServiceAnnotateResult{}
var retval *Communication
  var err2 error
  if retval, err2 = p.handler.Annotate(args.Original); err2 != nil {
  switch v := err2.(type) {
    case *ConcreteThriftException:
  result.Ex = v
    default:
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing annotate: " + err2.Error())
    oprot.WriteMessageBegin("annotate", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  }
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("annotate", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type annotateCommunicationServiceProcessorGetMetadata struct {
  handler AnnotateCommunicationService
}

func (p *annotateCommunicationServiceProcessorGetMetadata) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := AnnotateCommunicationServiceGetMetadataArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getMetadata", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := AnnotateCommunicationServiceGetMetadataResult{}
var retval *AnnotationMetadata
  var err2 error
  if retval, err2 = p.handler.GetMetadata(); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getMetadata: " + err2.Error())
    oprot.WriteMessageBegin("getMetadata", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("getMetadata", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type annotateCommunicationServiceProcessorGetDocumentation struct {
  handler AnnotateCommunicationService
}

func (p *annotateCommunicationServiceProcessorGetDocumentation) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := AnnotateCommunicationServiceGetDocumentationArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getDocumentation", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := AnnotateCommunicationServiceGetDocumentationResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.GetDocumentation(); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getDocumentation: " + err2.Error())
    oprot.WriteMessageBegin("getDocumentation", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("getDocumentation", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type annotateCommunicationServiceProcessorShutdown struct {
  handler AnnotateCommunicationService
}

func (p *annotateCommunicationServiceProcessorShutdown) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := AnnotateCommunicationServiceShutdownArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    return false, err
  }

  iprot.ReadMessageEnd()
  var err2 error
  if err2 = p.handler.Shutdown(); err2 != nil {
    return true, err2
  }
  return true, nil
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Original
type AnnotateCommunicationServiceAnnotateArgs struct {
  Original *Communication `thrift:"original,1" db:"original" json:"original"`
}

func NewAnnotateCommunicationServiceAnnotateArgs() *AnnotateCommunicationServiceAnnotateArgs {
  return &AnnotateCommunicationServiceAnnotateArgs{}
}

var AnnotateCommunicationServiceAnnotateArgs_Original_DEFAULT *Communication
func (p *AnnotateCommunicationServiceAnnotateArgs) GetOriginal() *Communication {
  if !p.IsSetOriginal() {
    return AnnotateCommunicationServiceAnnotateArgs_Original_DEFAULT
  }
return p.Original
}
func (p *AnnotateCommunicationServiceAnnotateArgs) IsSetOriginal() bool {
  return p.Original != nil
}

func (p *AnnotateCommunicationServiceAnnotateArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceAnnotateArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.Original = &Communication{}
  if err := p.Original.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Original), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceAnnotateArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("annotate_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AnnotateCommunicationServiceAnnotateArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("original", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:original: ", p), err) }
  if err := p.Original.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Original), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:original: ", p), err) }
  return err
}

func (p *AnnotateCommunicationServiceAnnotateArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnnotateCommunicationServiceAnnotateArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Ex
type AnnotateCommunicationServiceAnnotateResult struct {
  Success *Communication `thrift:"success,0" db:"success" json:"success,omitempty"`
  Ex *ConcreteThriftException `thrift:"ex,1" db:"ex" json:"ex,omitempty"`
}

func NewAnnotateCommunicationServiceAnnotateResult() *AnnotateCommunicationServiceAnnotateResult {
  return &AnnotateCommunicationServiceAnnotateResult{}
}

var AnnotateCommunicationServiceAnnotateResult_Success_DEFAULT *Communication
func (p *AnnotateCommunicationServiceAnnotateResult) GetSuccess() *Communication {
  if !p.IsSetSuccess() {
    return AnnotateCommunicationServiceAnnotateResult_Success_DEFAULT
  }
return p.Success
}
var AnnotateCommunicationServiceAnnotateResult_Ex_DEFAULT *ConcreteThriftException
func (p *AnnotateCommunicationServiceAnnotateResult) GetEx() *ConcreteThriftException {
  if !p.IsSetEx() {
    return AnnotateCommunicationServiceAnnotateResult_Ex_DEFAULT
  }
return p.Ex
}
func (p *AnnotateCommunicationServiceAnnotateResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *AnnotateCommunicationServiceAnnotateResult) IsSetEx() bool {
  return p.Ex != nil
}

func (p *AnnotateCommunicationServiceAnnotateResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceAnnotateResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &Communication{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceAnnotateResult)  ReadField1(iprot thrift.TProtocol) error {
  p.Ex = &ConcreteThriftException{}
  if err := p.Ex.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Ex), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceAnnotateResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("annotate_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AnnotateCommunicationServiceAnnotateResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *AnnotateCommunicationServiceAnnotateResult) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetEx() {
    if err := oprot.WriteFieldBegin("ex", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ex: ", p), err) }
    if err := p.Ex.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Ex), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ex: ", p), err) }
  }
  return err
}

func (p *AnnotateCommunicationServiceAnnotateResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnnotateCommunicationServiceAnnotateResult(%+v)", *p)
}

type AnnotateCommunicationServiceGetMetadataArgs struct {
}

func NewAnnotateCommunicationServiceGetMetadataArgs() *AnnotateCommunicationServiceGetMetadataArgs {
  return &AnnotateCommunicationServiceGetMetadataArgs{}
}

func (p *AnnotateCommunicationServiceGetMetadataArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceGetMetadataArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getMetadata_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AnnotateCommunicationServiceGetMetadataArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnnotateCommunicationServiceGetMetadataArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AnnotateCommunicationServiceGetMetadataResult struct {
  Success *AnnotationMetadata `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewAnnotateCommunicationServiceGetMetadataResult() *AnnotateCommunicationServiceGetMetadataResult {
  return &AnnotateCommunicationServiceGetMetadataResult{}
}

var AnnotateCommunicationServiceGetMetadataResult_Success_DEFAULT *AnnotationMetadata
func (p *AnnotateCommunicationServiceGetMetadataResult) GetSuccess() *AnnotationMetadata {
  if !p.IsSetSuccess() {
    return AnnotateCommunicationServiceGetMetadataResult_Success_DEFAULT
  }
return p.Success
}
func (p *AnnotateCommunicationServiceGetMetadataResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *AnnotateCommunicationServiceGetMetadataResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceGetMetadataResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &AnnotationMetadata{
  KBest: 1,
}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceGetMetadataResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getMetadata_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AnnotateCommunicationServiceGetMetadataResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *AnnotateCommunicationServiceGetMetadataResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnnotateCommunicationServiceGetMetadataResult(%+v)", *p)
}

type AnnotateCommunicationServiceGetDocumentationArgs struct {
}

func NewAnnotateCommunicationServiceGetDocumentationArgs() *AnnotateCommunicationServiceGetDocumentationArgs {
  return &AnnotateCommunicationServiceGetDocumentationArgs{}
}

func (p *AnnotateCommunicationServiceGetDocumentationArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceGetDocumentationArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getDocumentation_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AnnotateCommunicationServiceGetDocumentationArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnnotateCommunicationServiceGetDocumentationArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AnnotateCommunicationServiceGetDocumentationResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewAnnotateCommunicationServiceGetDocumentationResult() *AnnotateCommunicationServiceGetDocumentationResult {
  return &AnnotateCommunicationServiceGetDocumentationResult{}
}

var AnnotateCommunicationServiceGetDocumentationResult_Success_DEFAULT string
func (p *AnnotateCommunicationServiceGetDocumentationResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return AnnotateCommunicationServiceGetDocumentationResult_Success_DEFAULT
  }
return *p.Success
}
func (p *AnnotateCommunicationServiceGetDocumentationResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *AnnotateCommunicationServiceGetDocumentationResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceGetDocumentationResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *AnnotateCommunicationServiceGetDocumentationResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getDocumentation_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AnnotateCommunicationServiceGetDocumentationResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *AnnotateCommunicationServiceGetDocumentationResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnnotateCommunicationServiceGetDocumentationResult(%+v)", *p)
}

type AnnotateCommunicationServiceShutdownArgs struct {
}

func NewAnnotateCommunicationServiceShutdownArgs() *AnnotateCommunicationServiceShutdownArgs {
  return &AnnotateCommunicationServiceShutdownArgs{}
}

func (p *AnnotateCommunicationServiceShutdownArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *AnnotateCommunicationServiceShutdownArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("shutdown_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AnnotateCommunicationServiceShutdownArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AnnotateCommunicationServiceShutdownArgs(%+v)", *p)
}

