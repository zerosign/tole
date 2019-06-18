package engine

//
// This can be used for defining which
// template that need to evaluated if there is value changes.
//
type Relation struct {
	// template -> vars
	inner map[string]StrSet
	// vars -> template
	inverted map[string]StrSet
}

func EmptyRelation() Relation {
	return Relation{
		EmptyStrSet(),
		EmptyStrSet(),
	}
}

func (d *Relation) Relate(source str, target str) {

	var sets StrSet
	var ok bool
	if sets, ok = d.inner[source]; ok {
		sets.Add(target)
	} else {
		sets = NewStrSet()
		sets.Add(target)
	}
	d.inner[source] = sets
	d.keys = append(d.keys, source)

	if sets, ok = d.inverted[target]; ok {
		sets.Add(source)
	} else {
		sets = NewStrSet()
		sets.Add(source)
	}
	d.inverted[target] = sets
	d.invertedKeys = append(d.invertedKeys, target)
}

func (d *Relation) ImpactedSource(value str) []string {
	return d.inverted[value].Values()
}

func (d *Relation) ImpactedTarget(source str) []string {
	return d.inner[source].Values()
}
