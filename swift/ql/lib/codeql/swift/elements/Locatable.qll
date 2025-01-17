private import codeql.swift.generated.Locatable

class Locatable extends LocatableBase {
  pragma[nomagic]
  override Location getLocation() {
    result = LocatableBase.super.getLocation()
    or
    not exists(LocatableBase.super.getLocation()) and result instanceof UnknownLocation
  }
}
