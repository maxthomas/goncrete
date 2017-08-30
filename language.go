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

// A theory about what languages are present in a given communication
// or piece of communication.  Note that it is possible to have more
// than one language present in a given communication.
// 
// Attributes:
//  - UUID: Unique identifier for this language identification.
//  - Metadata: Information about where this language identification came from.
//  - LanguageToProbabilityMap: A list mapping from a language to the probability that that
// language occurs in a given communication.  Each language code should
// occur at most once in this list.  The probabilities do <i>not</i>
// need to sum to one -- for example, if a single communication is known
// to contain both English and French, then it would be appropriate
// to assign a probability of 1 to both langauges.  (Manually
// annotated LanguageProb objects should always have probabilities
// of either zero or one; machine-generated LanguageProbs may have
// intermediate probabilities.)
// 
// Note: The string key should represent the ISO 639-3 three-letter code.
type LanguageIdentification struct {
  UUID *UUID `thrift:"uuid,1,required" db:"uuid" json:"uuid"`
  Metadata *AnnotationMetadata `thrift:"metadata,2,required" db:"metadata" json:"metadata"`
  LanguageToProbabilityMap map[string]float64 `thrift:"languageToProbabilityMap,3,required" db:"languageToProbabilityMap" json:"languageToProbabilityMap"`
}

func NewLanguageIdentification() *LanguageIdentification {
  return &LanguageIdentification{}
}

var LanguageIdentification_UUID_DEFAULT *UUID
func (p *LanguageIdentification) GetUUID() *UUID {
  if !p.IsSetUUID() {
    return LanguageIdentification_UUID_DEFAULT
  }
return p.UUID
}
var LanguageIdentification_Metadata_DEFAULT *AnnotationMetadata
func (p *LanguageIdentification) GetMetadata() *AnnotationMetadata {
  if !p.IsSetMetadata() {
    return LanguageIdentification_Metadata_DEFAULT
  }
return p.Metadata
}

func (p *LanguageIdentification) GetLanguageToProbabilityMap() map[string]float64 {
  return p.LanguageToProbabilityMap
}
func (p *LanguageIdentification) IsSetUUID() bool {
  return p.UUID != nil
}

func (p *LanguageIdentification) IsSetMetadata() bool {
  return p.Metadata != nil
}

func (p *LanguageIdentification) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetUUID bool = false;
  var issetMetadata bool = false;
  var issetLanguageToProbabilityMap bool = false;

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
      issetUUID = true
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
      issetMetadata = true
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
      issetLanguageToProbabilityMap = true
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
  if !issetUUID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field UUID is not set"));
  }
  if !issetMetadata{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Metadata is not set"));
  }
  if !issetLanguageToProbabilityMap{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field LanguageToProbabilityMap is not set"));
  }
  return nil
}

func (p *LanguageIdentification)  ReadField1(iprot thrift.TProtocol) error {
  p.UUID = &UUID{}
  if err := p.UUID.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UUID), err)
  }
  return nil
}

func (p *LanguageIdentification)  ReadField2(iprot thrift.TProtocol) error {
  p.Metadata = &AnnotationMetadata{
  KBest: 1,
}
  if err := p.Metadata.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Metadata), err)
  }
  return nil
}

func (p *LanguageIdentification)  ReadField3(iprot thrift.TProtocol) error {
  _, _, size, err := iprot.ReadMapBegin()
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(map[string]float64, size)
  p.LanguageToProbabilityMap =  tMap
  for i := 0; i < size; i ++ {
var _key0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key0 = v
}
var _val1 float64
    if v, err := iprot.ReadDouble(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val1 = v
}
    p.LanguageToProbabilityMap[_key0] = _val1
  }
  if err := iprot.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *LanguageIdentification) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("LanguageIdentification"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *LanguageIdentification) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("uuid", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:uuid: ", p), err) }
  if err := p.UUID.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UUID), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:uuid: ", p), err) }
  return err
}

func (p *LanguageIdentification) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("metadata", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:metadata: ", p), err) }
  if err := p.Metadata.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Metadata), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:metadata: ", p), err) }
  return err
}

func (p *LanguageIdentification) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("languageToProbabilityMap", thrift.MAP, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:languageToProbabilityMap: ", p), err) }
  if err := oprot.WriteMapBegin(thrift.STRING, thrift.DOUBLE, len(p.LanguageToProbabilityMap)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
  }
  for k, v := range p.LanguageToProbabilityMap {
    if err := oprot.WriteString(string(k)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    if err := oprot.WriteDouble(float64(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteMapEnd(); err != nil {
    return thrift.PrependError("error writing map end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:languageToProbabilityMap: ", p), err) }
  return err
}

func (p *LanguageIdentification) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LanguageIdentification(%+v)", *p)
}
