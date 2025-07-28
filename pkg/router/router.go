package routerPkg

import "strings"

type HTTPFunc func(ctx any)
type params map[string]string

type Route struct {
	part       string
	children   []*Route
	handler    HTTPFunc
	isWildcard bool
	isParam    bool
}

func newRootNode() *Route {
	//root Route
	return &Route{
		part:     "/",
		children: make([]*Route, 0),
	}
}

func (n *Route) addChild(part string, isWildcard bool) *Route {
	child := &Route{
		part:       part,
		isWildcard: isWildcard,
		children:   make([]*Route, 0),
	}
	n.children = append(n.children, child)
	return child
}

func (n *Route) findChild(part string) *Route {
	for _, child := range n.children {
		if child.part == part {
			return child
		}
	}
	return nil
}

func (n *Route) findWildcardChild() *Route {
	for _, child := range n.children {
		if child.isWildcard {
			return child
		}
	}
	return nil
}

type Router struct {
	root *Route
}

func NewRouter() *Router {
	return &Router{
		root: newRootNode(),
	}
}

func (r *Router) AddRoute(path string, handler HTTPFunc) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	currentNode := r.root

	for _, part := range parts {
		isWildcard := strings.HasPrefix(part, ":")

		childNode := currentNode.findChild(part)

		if childNode == nil {
			childNode = currentNode.findWildcardChild()
			if childNode == nil || childNode.part != part {
				childNode = currentNode.addChild(part, isWildcard)
			}
		}
		currentNode = childNode
	}
	currentNode.handler = handler
}

func (r *Router) Match(path string) (HTTPFunc, params) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	currentNode := r.root
	extractedParams := make(params)

	for _, part := range parts {
		foundChild := false
		for _, child := range currentNode.children {
			if child.part == part {
				currentNode = child
				foundChild = true
				break
			}
		}

		if !foundChild {
			for _, child := range currentNode.children {
				if child.isWildcard {
					extractedParams[child.part[1:]] = part
					currentNode = child
					foundChild = true
					break
				}
			}
		}
		if !foundChild {
			return nil, nil
		}
	}

	if currentNode != nil && currentNode.handler != nil {
		return currentNode.handler, extractedParams
	}

	return nil, nil
}
